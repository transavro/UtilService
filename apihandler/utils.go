package apihandler

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type TileSchema struct {
	Content struct {
		Target       []string `json:"target"`
	} `json:"content"`
}

type DeadLinkCheckerResponse struct {
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		ID      string `json:"id"`
		Snippet struct {
			PlaylistID   string `json:"playlistId"`
			ResourceID   struct {
				Kind    string `json:"kind"`
				VideoID string `json:"videoId"`
			} `json:"resourceId"`
		} `json:"snippet,omitempty"`
	} `json:"items"`
}

//youtube dead link methods

var (
	youtubeApiKeys = [...]string{"AIzaSyCKUyMUlRTHMG9LFSXPYEDQYn7BCfjFQyI", "AIzaSyCNGkNspHPreQQPdT-q8KfQznq4S2YqjgU", "AIzaSyABJehNy0EEzzKl-I7hXkvYeRwIupl2RYA" }
	youtubeApiIndex = 0
	currentApiKeyUsed = youtubeApiKeys[youtubeApiIndex]
	deadFile *os.File;
	currentPlaylistCheckingId = ""
)

func parsingCollection(tilesCollection *mongo.Collection){
	log.Println("parsing collection....")
	myStages := mongo.Pipeline{
		bson.D{{"$match",  bson.D{{"content.package", "com.google.android.youtube"}}}},
		bson.D{{"$match",  bson.D{{"content.publishState", true}}}},
		bson.D{{"$project", bson.D{{"_id", 0}, {"content.target", 1}}}},
	}


	cur, err := tilesCollection.Aggregate(context.Background(), myStages)
	if err != nil {
		log.Fatal("Got Error", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("started parsing cursor")

	for cur.Next(context.Background()){
		var tile TileSchema
		err = cur.Decode(&tile)
		if err != nil {
			log.Fatalf("got error while decoding cursor obj to TileSchema %s", err)
		}

		if len(tile.Content.Target) == 0 {
			log.Fatalf("target array is empty")
		}

		for _, youtubeId := range tile.Content.Target{
			log.Println("checking youtubeId "+youtubeId)
			deadCheckerOptimus(youtubeId)
		}
	}

	cur.Close(context.Background())
}

func deadCheckerOptimus(youtubeId string){
	log.Println("deadCheckerOptimus ")
	youtubeEndpoint := "https://www.googleapis.com/youtube/v3/videos"
	isContentPlaylist := false


	if (strings.Compare(youtubeId[0:2], "PL") == 0) {
		log.Println("got playlist ")
		youtubeEndpoint = "https://www.googleapis.com/youtube/v3/playlistItems"
		isContentPlaylist = true
	}



	respReader := hitingYoutubeApi(youtubeEndpoint, youtubeId, currentApiKeyUsed, isContentPlaylist)
	log.Println("got io reader")
	if respReader != nil {
		parsingContent(respReader, isContentPlaylist, false, youtubeId)
	}
}

func hitingYoutubeApi(ytEndpoint, ytVideoId, ytApiKey string, isContentPlaylist bool ) (io.ReadCloser){
	log.Println("hittting api... ", ytEndpoint, " isContentPlayList ", isContentPlaylist)
	req, err := http.NewRequest("GET", ytEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	// here we define the query parameters and their respective values
	q := req.URL.Query()
	// notice how I'm using os.Getenv() to pick up the environment
	// variables that we defined earlier. No hard coded credentials here
	q.Add("key", ytApiKey)
	if isContentPlaylist {
		log.Println("playlist content query")
		q.Add("playlistId", ytVideoId)
	}else {
		log.Println("non playlist content query ")
		q.Add("id", ytVideoId)
	}
	q.Add("part", "snippet")
	req.URL.RawQuery = q.Encode()

	// finally we make the request to the URL that we have just
	// constructed
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("got server response ",resp.StatusCode)

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return nil
		}
		youtubeApiIndex ++ ;
		if youtubeApiIndex >= len(youtubeApiKeys) {
			log.Fatalf("All the youtube keys are used up . Please try tommorow . Press control + c")
		}else {
			currentApiKeyUsed = youtubeApiKeys[youtubeApiIndex]
			log.Println("Changing Api Keys ******************************************************************")
			result :=  hitingYoutubeApi(ytEndpoint, ytVideoId,  currentApiKeyUsed, isContentPlaylist)
			if result != nil {
				return result
			}
			return nil
		}
	}
	return resp.Body
}

func parsingContent(respReader io.ReadCloser, isContentPlaylist, isContentFromPlaylist bool, youtubeId string){
	log.Println("parsing content ")
	var response DeadLinkCheckerResponse
	body,err := ioutil.ReadAll(respReader)
	if err != nil {
		log.Fatal(err)
	}
	// and finally unmarshal it into an Response struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		respReader.Close()
		log.Fatalf("error while unmarshaling the json of youtube response ", err)
	}
	log.Println("marshaling or youtube response done ",response.PageInfo.ResultsPerPage," ",response.PageInfo.TotalResults)

	if(response.PageInfo.ResultsPerPage == 0 && response.PageInfo.TotalResults == 0){
		log.Printf("got dead link %s ",youtubeId)
		if isContentFromPlaylist {
			entryOfDeadInFile(youtubeId , currentPlaylistCheckingId)
		}else {
			entryOfDeadInFile(youtubeId, "")
		}
	}else if isContentPlaylist {
		log.Println("parsing playlist...", len(response.Items))
		currentPlaylistCheckingId = youtubeId
		for _, item := range response.Items {
			log.Println("playlist item ",item.Snippet.ResourceID.VideoID)
			respBody := hitingYoutubeApi("https://www.googleapis.com/youtube/v3/videos", item.Snippet.ResourceID.VideoID, currentApiKeyUsed, false)
			if respBody != nil {
				parsingContent(respBody, false, true, item.Snippet.ResourceID.VideoID)
			}
		}
	}
	respReader.Close()
}

func entryOfDeadInFile(youtubeVideoId ,youtubePlayListId  string){
	if len(youtubePlayListId) != 0 {
		_, err := fmt.Fprintln(deadFile, fmt.Sprintf("%s        %s",youtubePlayListId, youtubeVideoId))
		if err != nil {
			log.Fatalf("error while writing dead content in file %s", err)
		}
	}else {
		_, err := fmt.Fprintln(deadFile, fmt.Sprintf("%s", youtubeVideoId))
		if err != nil {
			log.Fatalf("error while writing dead content in file %s", err)
		}
	}
}
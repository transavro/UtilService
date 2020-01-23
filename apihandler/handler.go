package apihandler;

import (
	"context"
	pb "github.com/transavro/UtilService/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"strings"
)

type Server struct {
	TileCollection *mongo.Collection
}


func (s Server) DeadLinkChecker(ctx context.Context, target *pb.DeadLinkTarget) (*pb.DeadLinkResp, error){
	if strings.ToLower(target.GetSource()) == "youtube"{
		preTask()
		parsingCollection(s.TileCollection)
	}
	return &pb.DeadLinkResp{IsSuccessful:true}, nil
}

func preTask(){
	log.Println("init")
	os.Remove("dead.txt")
	deadFile , _ = os.Create("dead.txt")
	log.Println("init done")
}












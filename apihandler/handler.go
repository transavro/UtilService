package apihandler;

import (
	"fmt"
	pb "github.com/transavro/UtilService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strings"
)

type Server struct {
	TileCollection *mongo.Collection
	DeepLinkCollection *mongo.Collection
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

func(s Server) GetDeepLink(ctx context.Context, req *pb.DeepLinkReq) (*pb.DeepLinkResp, error){
	findResult := s.DeepLinkCollection.FindOne(ctx, bson.D{{"$and", bson.A{bson.D{{"packagename", req.GetPackageName()}}, bson.D{{"board", req.GetBoard()}}}}})
	if findResult.Err() != nil {
		return  nil, status.Error(codes.InvalidArgument, fmt.Sprintf("No deep link found for board %s and packageName %s ", req.GetBoard(), req.GetPackageName()))
	}
	var response pb.DeepLinkResp
	err := findResult.Decode(&response)
	if err != nil {
		return  nil, status.Error(codes.Internal, "Error while decoding to data")
	}
	return &response, nil
}












package main

import (
	"fmt"
	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/transavro/UtilService/apihandler"
	pb "github.com/transavro/UtilService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)



const (
	//atlasMongoHost          = "mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
	developmentMongoHost = "mongodb://dev-uni.cloudwalker.tv:6592"
	RedisHost   = ":6379"
	grpc_port        = ":7761"
	rest_port		 = ":7762"
)

// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota
)

var tileCollection *mongo.Collection

// Multiple init() function
func init() {
	fmt.Println("Welcome to init() function")
	tileCollection = getMongoCollection("cwtx2devel", "tiles", developmentMongoHost)
}

func credMatcher(headerName string) (mdName string, ok bool) {
	if headerName == "Login" || headerName == "Password" {
		return headerName, true
	}
	return "", false
}

// authenticateAgent check the client credentials
func authenticateClient(ctx context.Context, s *apihandler.Server) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		clientLogin := strings.Join(md["login"], "")
		clientPassword := strings.Join(md["password"], "")
		if clientLogin != "nayan" {
			return "", fmt.Errorf("unknown user %s", clientLogin)
		}
		if clientPassword != "makasare" {
			return "", fmt.Errorf("bad password %s", clientPassword)
		}
		log.Printf("authenticated client: %s", clientLogin)
		return "42", nil
	}
	return "", fmt.Errorf("missing credentials")
}

//uthenticate client using jwt
func auth(ctx context.Context) error {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(
			codes.InvalidArgument,
			"missing context",
		)
	}

	authString, ok := meta["authorization"]
	if !ok {
		return status.Errorf(
			codes.Unauthenticated,
			"missing authorization",
		)
	}
	// validate token algo
	log.Println("found jwt token")
	jwtToken, err := jwt.Parse(
		authString[0],
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("bad signing method")
			}
			// additional validation goes here.
			return []byte("transavro"), nil
		},
	)

	if jwtToken.Valid {
		return nil
	}
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return status.Error(codes.Internal, "bad token")
}

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	s, ok := info.Server.(*apihandler.Server)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("unable to cast the server"))
	}
	clientID, err := authenticateClient(ctx, s)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, clientIDKey, clientID)
	return handler(ctx, req)

	// auth using jwt

	//if err := auth(ctx); err != nil {
	//	return nil, err
	//}
	//log.Println("authorization OK")
	//return handler(ctx, req)
}

// streamAuthIntercept intercepts to validate authorization
func streamAuthIntercept(
	server interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	// auth using jwt
	if err := auth(stream.Context()); err != nil {
		return err
	}
	log.Println("authorization OK")
	return handler(server, stream)
}

func startGRPCServer(address string) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	} // create a server instance
	s := apihandler.Server{
		tileCollection,
	}

	// attach the Ping service to the server
	grpcServer := grpc.NewServer()                    // attach the Ping service to the server
	pb.RegisterUtilServiceServer(grpcServer, &s) // start the server
	//log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}

func startRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(credMatcher),)

	opts := []grpc.DialOption{grpc.WithInsecure()} // Register ping

	err := pb.RegisterUtilServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}

	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}

func getMongoCollection(dbName, collectionName, mongoHost string) *mongo.Collection {
	// Register custom codecs for protobuf Timestamp and wrapper types
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoHost), options.Client().SetRegistry(reg))
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient.Database(dbName).Collection(collectionName)
}

func getRedisClient(redisHost string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}
	return client
}

func main() {

	//grpcAddress := fmt.Sprintf("%s:%d", "cloudwalker.services.tv", 7775)
	//restAddress := fmt.Sprintf("%s:%d", "cloudwalker.services.tv", 7776)

	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpc_port)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(rest_port, grpc_port)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// infinite loop
	//log.Printf("Entering infinite loop")
	select {}
}
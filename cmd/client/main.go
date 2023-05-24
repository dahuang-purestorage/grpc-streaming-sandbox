package main

import (
	"context"
	"io"
	"log"

	"github.com/dahuang-purestorage/grpc-streaming-sandbox/pkg/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create new gRPC server instance
	serverAddr := "localhost:8080"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// req := wrappers.StringValue{
	// 	Value: "dahuang@purestorage.com",
	// }

	client := api.NewAdderClient(conn)
	req := api.GetRequest{}
	stream, err := client.GetNumber(context.Background(), &req)
	if err != nil {
		logrus.Errorf("error getting number from client %w", err)
	}

	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(result)
	}
	// client := phonebookservice.NewApplianceServiceClient(conn)
	// resp, err := client.ListAppliancesByFilter(context.Background(), &req)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// for {
	// 	appliance, err := resp.Recv()
	// 	fmt.Printf("resp %v ... %v error\n", appliance, err)
	// 	if err != nil {
	// 		break
	// 	}
	// }
}

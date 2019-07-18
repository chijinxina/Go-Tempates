package main

import (
	"context"
	pb "github.com/chijinxina/Go-Tempates/grpc/bidirectional-streaming/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	addr := "127.0.0.1:50070"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewUserServiceClient(conn)
	stream, err := grpcClient.GetUserInfo(context.Background())
	if err != nil {
		log.Fatalf("receive stream error: %v", err)
	}

	//向服务端发送数据流，并处理响应流
	var i int32
	for i = 1; i < 4; i++ {
		stream.Send(&pb.UserRequest{ID: i})
		time.Sleep(1 * time.Second)
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("response error: %v", err)
		}
		log.Printf("[RECEIVED RESPONSE]: %v\n", res)
	}
}

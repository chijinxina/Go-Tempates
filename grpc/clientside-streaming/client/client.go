package main

import (
	"context"
	pb "github.com/chijinxina/Go-Tempates/grpc/clientside-streaming/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	addr := "127.0.0.1:50060"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect to server error: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewUserServiceClient(conn)

	//向服务器发送流数据
	stream, err := grpcClient.GetUserInfo(context.Background())

	var i int32

	for i = 1; i < 4; i++ {
		err := stream.Send(&pb.UserRequest{ID: i})
		if err != nil {
			log.Fatalf("send error: %v", err)
		}
	}

	//接收服务端响应
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("received response error: %v", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", res)
}

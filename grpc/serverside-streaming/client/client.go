package main

import (
	"context"
	pb "github.com/chijinxina/Go-Tempates/grpc/serverside-streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	addr := "127.0.0.1:50050"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}
	defer conn.Close()

	//创建grpc客户端实例
	grpcClient := pb.NewUserServiceClient(conn)

	//客户端调用服务端函数
	req := pb.UserRequest{ID: 1}
	stream, err := grpcClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("receive response error: %v", err)
	}

	//接收流数据
	for {
		res, err := stream.Recv()
		if err == io.EOF { //服务端数据发送完毕
			break
		}
		if err != nil {
			log.Fatalf("receive error: %v", err)
		}
		log.Printf("[RECEIVED RESPONSE]: %v\n", res)
	}
}

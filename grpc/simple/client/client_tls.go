package main

import (
	"context"
	pb "github.com/chijinxina/Go-Tempates/grpc/simple/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	//grpc服务器地址
	addrWithTLS := "127.0.0.1:50052"

	//TLS连接
	creds, err := credentials.NewClientTLSFromFile("./grpc/simple/keys/server.pem", "macos")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	//使用TLS认证建立连接
	conn, err := grpc.Dial(addrWithTLS, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}

	defer conn.Close()

	//创建grpc客户端instance
	grpcClient := pb.NewUserServiceClient(conn)

	//调用服务端函数
	req := pb.UserRequest{ID: 1}
	res, err := grpcClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("received response error: %v", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", res)
}

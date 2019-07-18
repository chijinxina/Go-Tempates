package main

import (
	"context"
	pb "github.com/chijinxina/Go-Tempates/grpc/simple/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//grpc服务器地址
	addr := "127.0.0.1:8081"

	//不使用认证建立链接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
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

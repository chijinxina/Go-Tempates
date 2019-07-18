package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/chijinxina/Go-Tempates/grpc/simple/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

//模拟数据库查询结果
var usersWithToken = map[int32]pb.UserResponse{
	1: {Name: "chijinxin", Age: 25},
	2: {Name: "kenthonaa", Age: 11},
	3: {Name: "nmzasdasd", Age: 32},
}

type simpleServerWithToken struct{}

//simpleServer实现了pb中的UserServiceServer接口
func (s *simpleServerWithToken) GetUserInfo(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	//解析metadata中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("无Token认证信息")
	}

	var (
		appid  string
		appkey string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appid != "101010" || appkey != "i am key" {
		return nil, errors.New(fmt.Sprintf("Token认证无效： appid=%s, appkey=%s", appid, appkey))
	}

	var res *pb.UserResponse
	if user, ok := usersWithToken[req.ID]; ok {
		res = &user
	}
	log.Printf("[RECEIVED RESQUEST]: %v\n", req)
	return res, nil
}

func main() {
	//服务器监听
	addr := "0.0.0.0:50053"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Printf("Server listen on: %v with TLS + Token", addr)
	}

	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("./grpc/simple/keys/server.pem", "./grpc/simple/keys/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	//创建grpc服务器instance，并开启TLS认证
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//向grpc服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &simpleServerWithToken{})

	//启动grpc服务器阻塞等待客户端调用
	grpcServer.Serve(lis)
}

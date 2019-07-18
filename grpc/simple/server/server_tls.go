package main

import (
	"context"
	pb "github.com/chijinxina/Go-Tempates/grpc/simple/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

//模拟数据库查询结果
var usersWithTls = map[int32]pb.UserResponse{
	1: {Name: "chijinxin", Age: 25},
	2: {Name: "kenthonaa", Age: 11},
	3: {Name: "nmzasdasd", Age: 32},
}

type simpleServerWithTLS struct{}

//simpleServer实现了pb中的UserServiceServer接口
func (s *simpleServerWithTLS) GetUserInfo(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var res *pb.UserResponse
	if user, ok := usersWithTls[req.ID]; ok {
		res = &user
	}
	log.Printf("[RECEIVED RESQUEST]: %v\n", req)
	return res, nil
}

func main() {
	//服务器监听
	addr := "0.0.0.0:50052"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Printf("Server listen on: %v with TLS", addr)
	}

	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("./grpc/simple/keys/server.pem", "./grpc/simple/keys/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	//创建grpc服务器instance，并开启TLS认证
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//向grpc服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &simpleServerWithTLS{})

	//启动grpc服务器阻塞等待客户端调用
	grpcServer.Serve(lis)
}

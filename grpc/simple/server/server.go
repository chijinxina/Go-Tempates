package main

import (
	"context"
	pb "github.com/chijinxina/Go-Tempates/grpc/simple/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

//模拟数据库查询结果
var users = map[int32]pb.UserResponse{
	1: {Name: "chijinxin", Age: 25},
	2: {Name: "kenthonaa", Age: 11},
	3: {Name: "nmzasdasd", Age: 32},
}

type simpleServer struct{}

//simpleServer实现了pb中的UserServiceServer接口
func (s *simpleServer) GetUserInfo(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var res *pb.UserResponse
	if user, ok := users[req.ID]; ok {
		res = &user
	}
	log.Printf("[RECEIVED RESQUEST]: %v\n", req)
	return res, nil
}

func main() {
	addr := "0.0.0.0:8081"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Println("Server listen on: ", addr)
	}

	//创建grpc服务器instance
	grpcServer := grpc.NewServer()
	//向grpc服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &simpleServer{})

	//启动grpc服务器阻塞等待客户端调用
	grpcServer.Serve(lis)
}

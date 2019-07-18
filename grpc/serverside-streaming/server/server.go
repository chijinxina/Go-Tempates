package main

import (
	pb "github.com/chijinxina/Go-Tempates/grpc/serverside-streaming/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

//模拟数据库查询结果
var users = map[int32]pb.UserResponse{
	1: {Name: "chijinxin", Age: 25},
	2: {Name: "asdasfsdd", Age: 32},
	3: {Name: "ADSFASFDG", Age: 12},
}

type serverSideStreamServer struct{}

func (s *serverSideStreamServer) GetUserInfo(req *pb.UserRequest, stream pb.UserService_GetUserInfoServer) error {
	//响应流数据
	for i := 0; i < 10; i++ {
		for _, user := range users {
			if err := stream.Send(&user); err != nil {
				return err
			}
		}
	}

	log.Printf("[RECEIVED REQUEST]: %v\n", req)
	return nil
}

func main() {
	addr := "0.0.0.0:50050"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Println("server listen:", addr)
	}

	//创建grpc服务器实例
	grpcServer := grpc.NewServer()

	//向grpc服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &serverSideStreamServer{})

	//启动grpc服务器，阻塞等待客户端调用
	grpcServer.Serve(lis)
}

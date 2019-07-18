package main

import (
	pb "github.com/chijinxina/Go-Tempates/grpc/clientside-streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

//模拟数据库查询结果
var users = map[int32]pb.UserResponse{
	1: {Name: "chijinxin", Age: 25},
	2: {Name: "asdasfsdd", Age: 32},
	3: {Name: "ADSFASFDG", Age: 12},
}

type clientSideStreamServer struct{}

func (s *clientSideStreamServer) GetUserInfo(stream pb.UserService_GetUserInfoServer) error {
	var lastID int32

	for {
		req, err := stream.Recv()
		//客户端流数据发送完毕
		if err == io.EOF {
			//返回最后一个ID的用户信息
			if u, ok := users[lastID]; ok {
				stream.SendAndClose(&u)
				return nil
			}
		}
		lastID = req.ID
		log.Printf("[RECEIVED REQUEST]: %v\n", req)
	}
	return nil
}

func main() {
	addr := "0.0.0.0:50060"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Println("server listen on: ", addr)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &clientSideStreamServer{})

	grpcServer.Serve(lis)
}

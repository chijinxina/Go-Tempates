package main

import (
	pb "github.com/chijinxina/Go-Tempates/grpc/bidirectional-streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

// 模拟的数据库查询结果
var users = map[int32]pb.UserResponse{
	1: {Name: "Dennis MacAlistair Ritchie", Age: 70},
	2: {Name: "Ken Thompson", Age: 75},
	3: {Name: "Rob Pike", Age: 62},
}

type bidirectionalStreamServer struct{}

func (s *bidirectionalStreamServer) GetUserInfo(stream pb.UserService_GetUserInfoServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		u := users[req.ID]
		err = stream.Send(&u)
		if err != nil {
			return err
		}
		log.Printf("[RECEIVED REQUEST]: %v\n", req)
	}
	return nil
}

func main() {
	addr := "0.0.0.0:50070"
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	} else {
		log.Println("server listen: ", addr)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &bidirectionalStreamServer{})

	grpcServer.Serve(lis)
}

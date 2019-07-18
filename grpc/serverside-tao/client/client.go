package main

import (
	"context"
	"flag"
	pb "github.com/chijinxina/Go-Tempates/grpc/serverside-tao/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

var (
	agentName string
)

func init() {
	flag.StringVar(&agentName, "n", "agent1", "agent address")
}

func main() {
	flag.Parse()

	addr := "127.0.0.1:50050"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewMTaskServiceClient(conn)

	req := pb.MTaskRequest{AgentName: agentName}

	var stream pb.MTaskService_GetMTaskClient
	//var err error
	stream, err = grpcClient.GetMTask(context.Background(), &req)
	if err != nil {
		log.Fatalf("receive response error: %v", err)
	}

	//接收服务端流数据
	for {
		res, err1 := stream.Recv()
		if err1 == io.EOF {
			log.Println("recevice from grpc server complete")
			for i := 1; i < 100; i++ {
				stream, err = grpcClient.GetMTask(context.Background(), &pb.MTaskRequest{AgentName: agentName})
				if err != nil {
					time.Sleep(1 * time.Second)
					log.Println("try to recreate stream ......")
				} else {
					break
				}
			}
			log.Println("create stream success")
		}
		if err1 != nil && err1 != io.EOF {
			log.Fatalf("receive error: %v\n", err1)
		}
		log.Printf("[RECEIVED RESPONSE]: %v\n", res)
	}

}

package main

import (
	"errors"
	"fmt"
	pb "github.com/chijinxina/Go-Tempates/grpc/serverside-tao/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

//type grpcClientInfo struct {
//
//}

type MTaskMsgDispatcher struct {
	grpcClientSession sync.Map
}

//查询grpc会话是否存在
func (d MTaskMsgDispatcher) Query(agentAddr string) bool {
	if _, ok := d.grpcClientSession.Load(agentAddr); ok {
		return true
	} else {
		return false
	}
}

//添加grpc会话
func (d *MTaskMsgDispatcher) Add(agentAddr string, sendMsgChan chan *pb.MTaskResponse) error {
	if val, _ := d.grpcClientSession.LoadOrStore(agentAddr, sendMsgChan); val != nil {
		return nil
	} else {
		return errors.New("Add gRPC ClientSession to MTaskMsgDispatcher Failed!\n")
	}

}

//删除grpc会话
func (d *MTaskMsgDispatcher) Del(agentAddr string) {
	//先关闭通道
	if c, ok := d.grpcClientSession.Load(agentAddr); ok {
		close(c.(chan *pb.MTaskResponse))
	}
	d.grpcClientSession.Delete(agentAddr)
}

//向对应会话通道中写入数据
func (d *MTaskMsgDispatcher) Send(agentAddr string, data string) {
	if c, ok := d.grpcClientSession.Load(agentAddr); ok {
		s := &pb.MTaskResponse{TaskId: data}
		c.(chan *pb.MTaskResponse) <- s
	}
}

type MTaskTransServer struct {
	//grpc MTask传输消息分发器
	dispatcher *MTaskMsgDispatcher
	//缓存MTask信息队列长度
	maxQueueLen int32
}

func (s *MTaskTransServer) GetMTask(req *pb.MTaskRequest, stream pb.MTaskService_GetMTaskServer) error {
	log.Printf("Agent: %v join the grpc session\n", req.AgentName)
	sendMsgChan := make(chan *pb.MTaskResponse, s.maxQueueLen)
	//在MTask消息分发器中添加会话
	if err := s.dispatcher.Add(req.AgentName, sendMsgChan); err != nil {
		return err
	}
	//服务调用结束后删除grpc会话
	defer func() {
		s.dispatcher.Del(req.AgentName)
		log.Printf("session: ")
		s.dispatcher.grpcClientSession.Range(func(key, value interface{}) bool {
			fmt.Printf("%v ", key)
			return true
		})
		log.Printf("agent: %v is leaved\n", req.AgentName)
	}()

	//循环读取MTask消息分发器通道中的数据，有数据则通过stream发送至grpc客户端
	for {
		select {
		//客户端关闭了grpc stream
		case <-stream.Context().Done():
			log.Println("grpc server stream closed")
			return nil

		case res, ok := <-sendMsgChan:
			if ok {
				//读取到数据，则通过stream将数据发送出去
				if err := stream.Send(res); err != nil {
					return err
				}
			} else {
				//通道关闭，则退出服务
				log.Println("channel closed")
				break
			}

		case <-time.After(8 * time.Second):
			return nil
		}

	}
	return nil
}

func input(d *MTaskMsgDispatcher) {
	for {
		var client, s string
		fmt.Scan(&client, &s)
		fmt.Println("Input: ", client, s)
		d.Send(client, s)
	}
}

func main() {

	addr := "0.0.0.0:50050"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Println("server listen:", addr)
	}

	var ddd MTaskMsgDispatcher

	go input(&ddd)
	//创建grpc服务器实例
	grpcServer := grpc.NewServer()

	//向grpc服务器注册服务
	pb.RegisterMTaskServiceServer(grpcServer, &MTaskTransServer{dispatcher: &ddd, maxQueueLen: 100})

	//启动grpc服务器，阻塞等待客户端调用
	grpcServer.Serve(lis)
}

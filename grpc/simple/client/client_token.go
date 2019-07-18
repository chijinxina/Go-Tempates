package main

import (
	"context"
	pb "github.com/chijinxina/Go-Tempates/grpc/simple/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

//自定义认证
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "1010110",
		"appkey": "i am key",
	}, nil
}
func (c customCredential) RequireTransportSecurity() bool {
	return true
}

func main() {
	//grpc服务器地址
	addrWithToken := "127.0.0.1:50053"
	var opts []grpc.DialOption

	//TLS连接
	creds, err := credentials.NewClientTLSFromFile("./grpc/simple/keys/server.pem", "macos")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	//使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	//使用TLS认证建立连接
	conn, err := grpc.Dial(addrWithToken, opts...)
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}

	defer conn.Close()

	//创建grpc客户端instance
	grpcClient := pb.NewUserServiceClient(conn)

	//调用服务端函数
	req := pb.UserRequest{ID: 2}
	res, err := grpcClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("received response error: %v", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", res)
}

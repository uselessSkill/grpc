package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	pb "grpc/pb"
)

const (
	address     = ":50051"
	defaultName = "wangxiaoning"
)

func main() {
	//连接grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//创建 GreeterService 的客户端对象
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//发送 RPC 请求，等待同步响应，得到回调后返回响应结果
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}


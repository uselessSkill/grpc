package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "grpc/pb"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	//创建 Listen，监听 TCP 端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创建 gRPC Server对象
	s := grpc.NewServer()
	//将 GreeterServer（其包含需要被调用的服务端接口）注册到gRPC Server 的内部注册中心
	//这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	//gRPC Server开始 lis.Accept,直到 Stop 或 GracefulStop
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


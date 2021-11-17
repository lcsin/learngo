package main

import (
	"context"
	"fmt"
	hello "gihub.com/lcsin/learngo/cncf/gRPC/quickstart/pb"
	"google.golang.org/grpc"
	"net"
)

/*
1. 取出server
2. 挂载方法
3. 注册服务
4. 创建监听
*/
type server struct {
	hello.UnimplementedHelloGRPCServer
}

func (s *server) SayHello(ctx context.Context, req *hello.HelloRequest) (rep *hello.HelloReply, err error) {
	fmt.Println("request:", req)
	return &hello.HelloReply{Message: "hello,grpc!"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("cannot listen rpc, err:", err)
	}
	defer listen.Close()

	s := grpc.NewServer()
	hello.RegisterHelloGRPCServer(s, &server{})
	s.Serve(listen)
}

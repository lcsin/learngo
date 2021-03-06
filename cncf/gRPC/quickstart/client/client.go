package main

import (
	"context"
	"fmt"
	hello "gihub.com/lcsin/learngo/cncf/gRPC/quickstart/pb"
	"google.golang.org/grpc"
)

/*
1. 创建一个链接
2. new一个client
3. 调用client方法
4. 获取返回值
*/
func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println("create connection failed,err:", err)
	}
	defer conn.Close()

	client := hello.NewHelloGRPCClient(conn)
	reply, err := client.SayHello(context.Background(), &hello.HelloRequest{Message: "give me five"})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("reply:", reply.Message)
}

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "learngo/microservice/gRPC/quickstart/pb"
)

/*
1. 创建一个链接
2. new一个client
3. 调用client方法
4. 获取返回值
*/
func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
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

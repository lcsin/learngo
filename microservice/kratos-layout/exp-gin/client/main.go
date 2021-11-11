package main

import (
	"context"
	v1 "exp-gin/api/user/v1"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("create connection failed,err:", err)
	}
	defer conn.Close()

	client := v1.NewUserServiceClient(conn)
	users, err := client.ListUser(context.Background(), &v1.ListUserRequest{})
	if err != nil {
		fmt.Println("gRPC ListUser error:", err)
	}
	fmt.Println("users:", users)
}

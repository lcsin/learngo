/*
go操作etcd
@see https://www.liwenzhou.com/posts/Go/go_etcd/
*/
package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Conn(endpoint []string) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoint,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, err:", err)
		return nil, err
	}
	return cli, nil
}

func Put(ctx context.Context, client clientv3.Client, key, value string) (*clientv3.PutResponse, error) {
	r, err := client.Put(ctx, key, value)
	if err != nil {
		fmt.Println("put to etcd failed,err:", err)
		return nil, err
	}
	return r, nil
}

func get(ctx context.Context, client clientv3.Client, key string) (*clientv3.GetResponse, error) {
	r, err := client.Get(ctx, key)
	if err != nil {
		fmt.Println("get from etcd failed,err:", err)
		return nil, err
	}
	return r, nil
}

func main() {
	_, err := Conn([]string{"192.168.5.130:9527"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("connect success!")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler start")

	ctx := r.Context()
	select {
	case <-time.After(5 * time.Second): // 表示经过5秒
		fmt.Println("finishing doing something")
	case <-ctx.Done(): // 表示context被取消
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("handler end")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

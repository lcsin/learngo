package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Article struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID       uint
	Nickname string
	Password string
}

/*
1. 引入sqlx、mysql依赖
2. 使用sqlx.Open代替sql.Open
*/
func main() {

	db, err := sqlx.Open("mysql", "root:root@tcp(10.225.137.237:3306)/kratos")
	if err != nil {
		fmt.Printf("cannot opening connection to mysql:%v", err)
	}

	var user User
	err = db.Get(&user, "select nickname,password from tb_user where nickname = ?", "admin")
	if err != nil {
		fmt.Println("get user failed,err:", err)
	}
	fmt.Println("user:", user)

}

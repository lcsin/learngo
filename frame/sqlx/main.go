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

/*
1. 引入sqlx、mysql依赖
2. 使用sqlx.Open代替sql.Open
*/
func main() {

	db, err := sqlx.Open("mysql", "root:root@tcp(10.225.137.237:3306)/testdb")
	if err != nil {
		fmt.Printf("cannot opening connection to mysql:%v", err)
	}

	var article Article
	err = db.Get(&article, "select id, title from articles where id = ?", 1)
	if err != nil {
		fmt.Println("get article failed,err:", err)
	}
	fmt.Println("article:", article)

}

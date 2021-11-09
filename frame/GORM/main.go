package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"learngo/frame/GORM/models"
)

func main() {
	db1Dsb := "root:root@tcp(10.225.137.237:3306)/kf_system?charset=utf8&parseTime=True&loc=Local&timeout=1000ms"
	db2Dsb := "bfcs:9YQPBcOh@tcp(rm-bp12r1j5cx2cla7q1vo.mysql.rds.aliyuncs.com:3306)/kf_system?charset=utf8&parseTime=True&loc=Local&timeout=1000ms"

	db, err := gorm.Open(mysql.Open(db1Dsb), &gorm.Config{})
	if err != nil {
		fmt.Println("cannot connect db1")
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		// `db2` 作为 sources，`db3`、`db4` 作为 replicas
		Sources:  []gorm.Dialector{mysql.Open(db1Dsb)},
		Replicas: []gorm.Dialector{mysql.Open(db2Dsb)},
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))

	var users []models.KfUsers
	err = db.Model(&models.KfUsers{}).Find(&users).Error
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("users:", users)

	var group []models.KfGroup
	err = db.Table("group").Find(&group).Error
	if err != nil {
		fmt.Println("err:", err.Error(), db.Name())
	}
	fmt.Println("group:", group)

	var category []models.KfCategory
	err = db.Table("category").Find(&category).Error
	if err != nil {
		fmt.Println("err:", err.Error(), db.Name())
	}
	fmt.Println("category:", category)
}

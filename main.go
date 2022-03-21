package main

import (
	"fmt"
	"go-demos/db"
	"go-demos/route"
)

func main() {
	// 代码启动，先连接数据库
	err := db.ConnDB("./photo.sqlite")
	if err != nil {
		panic("数据库连接错误:" + err.Error())
	}

	fmt.Println("数据库连接成功")

	route.InitRoute()

}

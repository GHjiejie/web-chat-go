package main

import (
	// 导入pkg里面的包
	"fmt"
	"log"
	"net/http"
	"webChatGo/pkg/database"
	"webChatGo/pkg/router"
)

func main() {
	database.Connect()        //进行数据库的连接
	r := router.SetupRouter() //设置路由
	// 启动服务器
	log.Fatal(http.ListenAndServe(":3000", r)) // 启动服务器
	fmt.Println("服务在3000端口启动")

}

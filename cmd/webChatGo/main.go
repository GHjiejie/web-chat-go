package main

import (
	// 导入pkg里面的包
	"log"
	"net/http"
	"webChatGo/pkg/database"
	"webChatGo/pkg/router"
)

func main() {
	database.Connect()        //进行数据库的连接
	r := router.SetupRouter() //设置路由
	// 启动服务器
	log.Fatal(http.ListenAndServe(":8080", r)) // 启动服务器

}

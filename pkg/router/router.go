package router

import (
	"webChatGo/pkg/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// 用户注册路由
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	return router
}

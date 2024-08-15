// router/router.go
package router

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// 注册用户相关路由
	UserRoutes(router)

	// 注册聊天相关路由
	// ChatRoutes(router)

	return router
}

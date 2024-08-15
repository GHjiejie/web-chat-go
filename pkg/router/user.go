// router/user.go
package router

import (
	"webChatGo/pkg/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/v1/user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/v1/user/login", controllers.Login).Methods("POST")
}

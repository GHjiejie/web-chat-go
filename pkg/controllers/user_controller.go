package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"webChatGo/pkg/database"
	// 导入数据库包
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register handles user registration
// 第一个参数是一个http.ResponseWriter类型的指针，用于向客户端发送HTTP响应。
// 第二个参数是一个指向http.Request类型的指针，包含了客户端发送的HTTP请求。我们可以从这个参数中读取请求体。
func Register(w http.ResponseWriter, r *http.Request) {

	// 输出请求体

	// log.Println("请求体:", &r.Body)

	log.Println("成功匹配到用户注册路由")
	var req RegisterRequest
	// body, err := io.ReadAll(r.Body)
	// 解析请求体
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Println("解析请求体失败:", err)

		http.Error(w, "解析请求体失败了", http.StatusBadRequest)
		return
	} else {
		log.Println("解析请求体成功")
	}

	// 验证输入(不能为空)
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// TODO: 与数据库交互，保存用户信息

	// 创建用户实体
	user := database.User{
		Username: req.Username,
		Password: req.Password,
	}

	// 将用户信息插入数据库
	err = database.RegisterUser(database.DB, user)
	if err != nil {
		http.Error(w, "用户插入失败", http.StatusInternalServerError)
		return
	} else {
		log.Println("用户插入成功")
	}

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})

}

package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"webChatGo/pkg/database"
)

// 用户注册结构体
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户登录结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户注册·
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	log.Println("成功匹配到用户注册路由")

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
		http.Error(w, "用户名和密码不能为空", http.StatusBadRequest)
		return
	}

	// 创建用户实体
	user := database.User{
		Username: req.Username,
		Password: req.Password,
	}

	// 将用户信息插入数据库，调用数据库包中的 RegisterUser 函数
	err = database.RegisterUser(database.DB, user)
	if err != nil {
		log.Println("用户注册失败:", err)
		http.Error(w, "用户注册失败", http.StatusInternalServerError)
		return
	} else {
		log.Println("用户注册成功")
	}

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "用户注册成功"})
}

// 这个w和r是http包中的ResponseWriter和Request类型，分别代表响应和请求。我们可以使用这两个参数来读取请求体、写入响应体、设置响应状态码等。
func Login(w http.ResponseWriter, r *http.Request) {

	// 我们还是要使用json包来解析请求体

	var req LoginRequest
	log.Println("成功匹配到用户登录路由")

	// 解析请求体
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("解析请求体失败:", err)
		http.Error(w, "解析请求体失败了", http.StatusBadRequest)
		return
	} else {
		log.Println("解析请求体成功")
	}

	// 判断用户输入的数据是否为空
	// 我们前端传输的是username与password，但是为什么我们使用req.Username与req.Password呢？
	// 这是因为我们在定义结构体时，使用了json标签，这样就可以将前端传输的数据与结构体中的字段进行匹配。就类似于一个map映射
	if req.Username == "" || req.Password == "" {
		http.Error(w, "用户名和密码不能为空", http.StatusBadRequest)
		return
	}

	// 查询数据库里面是否存在该用户的消息
	// 解释一下为什么有两个参数哈，可能在查询的过程中发生错误，这样就会返回err,如果找到用户的消息，就会使用user这个变量来接受我们数据库查询成功的值
	user, err := database.GetUserByUsername(database.DB, req.Username)
	if err != nil {
		log.Println("查询用户失败:", err)
		http.Error(w, "用户不存在，请先注册", http.StatusInternalServerError)
		return
	}

	// 判断用户输入的密码是否正确
	if user.Password != req.Password {
		http.Error(w, "密码错误", http.StatusUnauthorized)
		return
	}

	// 准备返回的用户信息，排除密码
	response := map[string]interface{}{
		"message": "登录成功",
		"user": map[string]interface{}{
			"username": user.Username,
			"id":       user.ID,
			// 可以添加更多用户信息
		},
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

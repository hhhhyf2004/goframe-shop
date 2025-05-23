package login

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tag:"登录" summary:"登录"`
	Name     string `json:"name" v:"required#用户名不能为空" `
	Password string `json:"password" v:"required#密码不能为空" `
}

type LoginRes struct {
	UserKey string `json:"user_key"`
	Token   string `json:"token"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tag:"登录" summary:"登出"`
}

type LogoutRes struct{}

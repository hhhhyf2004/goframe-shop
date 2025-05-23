package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta   `path:"/admin/add" tags:"管理员" method:"post" summary:"创建"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色ids"`
	IsAdmin  int    `json:"is_admin"    dc:"是否超级管理员"`
}

type CreateRes struct {
	AdminId int `json:"admin_id"`
}

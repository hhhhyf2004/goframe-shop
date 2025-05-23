package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta   `path:"/admin/update/" method:"post" tags:"管理员" summary:"修改"`
	Id       uint   `json:"id"  v:"min:1#请选择需要修改的管理员" dc:"管理员Id"`
	Name     string `json:"name"  dc:"用户名"`
	Password string `json:"password"   dc:"密码"`
	RoleIds  string `json:"role_ids"   dc:"角色ids"`
	IsAdmin  int    `json:"is_admin"   dc:"是否超级管理员"`
}
type UpdateRes struct {
	Id uint `json:"id"`
}

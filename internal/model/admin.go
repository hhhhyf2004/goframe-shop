package model

import "github.com/gogf/gf/v2/os/gtime"

type AdminGetListItem struct {
	Id        uint        `json:"id"` // 自增ID
	Name      string      `json:"name"`
	RoleIds   string      `json:"role_ids"`
	IsAdmin   uint        `json:"is_admin"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

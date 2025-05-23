package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteReq struct {
	g.Meta `path:"/admin/delete" method:"delete" tags:"管理员" summary:"删除"`
	Id     uint `v:"min:1#请选择需要删除的管理员" dc:"管理员id"`
}
type DeleteRes struct{}

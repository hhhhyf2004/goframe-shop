package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteReq struct {
	g.Meta `path:"/backend/rotation/delete" method:"delete" tags:"轮播图" summary:"删除"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}

type DeleteRes struct{}

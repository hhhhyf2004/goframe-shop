package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta `path:"/rotation/update" method:"post" tags:"轮播图" summary:"修改"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	PicUrl string `json:"pic_url"  dc:"图片链接"`
	Link   string `json:"link"  dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"跳转链接"`
}
type UpdateRes struct {
	Id uint `json:"id"`
}

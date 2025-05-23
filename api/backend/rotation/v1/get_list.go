package v1

import (
	"goframe-shop/api/backend"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/backend/rotation/list" method:"get" tags:"轮播图" summary:"列表"`
	backend.CommonPaginationReq
}
type GetListRes struct {
	backend.CommonPaginationRes
}

package admin

import (
	"context"

	v1 "goframe-shop/api/backend/admin/v1"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	var list []model.AdminGetListItem
	err = dao.AdminInfo.Ctx(ctx).Page(req.Page, req.Size).ScanAndCount(&list, &res.Total, false)
	if err != nil {
		return nil, err 
	}
	res.List = list 
	return 
}

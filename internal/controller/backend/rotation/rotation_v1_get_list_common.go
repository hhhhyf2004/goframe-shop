package rotation

import (
	"context"

	v1 "goframe-shop/api/backend/rotation/v1"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	var list []model.RotationItem
	err = dao.RotationInfo.
		Ctx(ctx).
		Page(req.Page, req.Size).
		ScanAndCount(&list, &res.Total, false)
	res.List = list
	return
}

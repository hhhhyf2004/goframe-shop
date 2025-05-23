package admin

import (
	"context"

	v1 "goframe-shop/api/backend/admin/v1"
	"goframe-shop/internal/dao"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Id, req.Id).Delete()
	return 
}

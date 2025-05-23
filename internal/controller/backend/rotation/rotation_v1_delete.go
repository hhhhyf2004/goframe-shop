package rotation

import (
	"context"

	v1 "goframe-shop/api/backend/rotation/v1"
	"goframe-shop/internal/dao"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.RotationInfo.
		Ctx(ctx).
		Where(dao.RotationInfo.Columns().Id, req.Id).
		Unscoped().
		Delete()
	return 
}

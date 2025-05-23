package rotation

import (
	"context"

	v1 "goframe-shop/api/backend/rotation/v1"
	"goframe-shop/internal/dao"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	cls := dao.RotationInfo.Columns()
	_, err = dao.RotationInfo.
		Ctx(ctx).
		FieldsEx(cls.Id).
		Where(cls.Id, req.Id).
		Data(req).
		OmitEmpty().
		Update()
	return 
}

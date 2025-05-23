package rotation

import (
	"context"

	v1 "goframe-shop/api/backend/rotation/v1"
	"goframe-shop/internal/dao"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := dao.RotationInfo.
		Ctx(ctx).
		Data(req).
		InsertAndGetId()
	if err != nil {
		return nil, err 
	}
	return &v1.CreateRes{RotationId: int(id)}, nil 
}

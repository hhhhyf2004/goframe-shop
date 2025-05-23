package admin

import (
	"context"

	v1 "goframe-shop/api/backend/admin/v1"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model/entity"
	"goframe-shop/utility"

	"github.com/gogf/gf/v2/util/grand"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	admin := entity.AdminInfo{
		Name: req.Name,
		RoleIds: req.RoleIds,
		IsAdmin: req.IsAdmin,
	}
	admin.UserSalt = grand.S(10)
	admin.Password = utility.EncryptPassword(admin.Password, admin.UserSalt)
	id, err := dao.AdminInfo.
		Ctx(ctx).
		Data(admin).
		OmitEmpty().
		InsertAndGetId()
	return &v1.CreateRes{AdminId: int(id)}, err 
}

package admin

import (
	"context"

	v1 "goframe-shop/api/backend/admin/v1"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model/entity"
	"goframe-shop/utility"

	"github.com/gogf/gf/v2/util/grand"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	admin := entity.AdminInfo{
		Id:       int(req.Id),
		Name:     req.Name,
		IsAdmin:  req.IsAdmin,
		RoleIds:  req.RoleIds,
	}
	if req.Password != "" {
		admin.UserSalt = grand.S(10)
		admin.Password = utility.EncryptPassword(admin.Password, admin.UserSalt)
	}
	cls := dao.AdminInfo.Columns()
	_, err = dao.AdminInfo.
		Ctx(ctx).
		Data(admin).
		Where(cls.Id, admin.Id). 
		FieldsEx(cls.Id, admin.Id).
		OmitEmpty().
		Update()
	return 
}

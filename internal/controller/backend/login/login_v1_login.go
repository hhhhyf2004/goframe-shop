package login

import (
	"context"
	"strconv"

	v1 "goframe-shop/api/backend/login/v1"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model/entity"
	"goframe-shop/utility"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	admin := entity.AdminInfo{}
	dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, req.Name).Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin.Password != utility.EncryptPassword(req.Password, admin.UserSalt) {
		return nil, gerror.New("账号或密码不正确")
	}
	userKey := consts.BACKEND_USERKEY_PREFIX + strconv.Itoa(admin.Id)
	token, err := consts.GToken.Generate(ctx, userKey, nil)
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{
		UserKey: userKey,
		Token:   token,
	}, nil
}

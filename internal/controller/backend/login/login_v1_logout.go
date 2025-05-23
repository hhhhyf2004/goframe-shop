package login

import (
	"context"

	v1 "goframe-shop/api/backend/login/v1"
	"goframe-shop/internal/consts"

	"github.com/goflyfox/gtoken/v2/gtoken"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	err = consts.GToken.Destroy(ctx, ctx.Value(gtoken.KeyUserKey).(string))
	return 
}

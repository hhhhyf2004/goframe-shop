package cmd

import (
	"context"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/controller/backend/admin"
	"goframe-shop/internal/controller/backend/login"
	"goframe-shop/internal/controller/backend/rotation"

	"github.com/goflyfox/gtoken/v2/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			options := &gtoken.Options{}
			err = g.Cfg().MustGet(ctx, "gToken").Struct(options)
			if err != nil {
				panic("options init fail")
			}
			// 创建gtoken对象
			consts.GToken = gtoken.NewDefaultToken(*options)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(CORS)
				group.Bind(login.NewV1().Login)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(gtoken.NewDefaultMiddleware(consts.GToken).Auth)
					group.Bind(
						rotation.NewV1(),
						admin.NewV1(),
						login.NewV1().Logout,
					) 
				})
			})
			s.Run()
			return nil
		},
	}
)

// 跨域
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

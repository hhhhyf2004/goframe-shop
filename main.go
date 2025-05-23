package main

import (
	_ "goframe-shop/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-shop/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

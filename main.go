package main

import (
	_ "UniApi/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"UniApi/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

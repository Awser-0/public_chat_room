package main

import (
	_ "chat/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"chat/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

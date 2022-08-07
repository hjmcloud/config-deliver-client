package main

import (
	_ "config-deliver-client/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"config-deliver-client/internal/cmd"
)

func main() {
	// cmd.Main.Run(gctx.New())
	cmd.Task.Run(gctx.New())
}

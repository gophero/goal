package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

const (
	Y = "y"
	N = "n"
)

func Confirm(ctx *cli.Context, title string, fn func()) {
	skipConfirm := ctx.Bool("q")
	if skipConfirm {
		fn()
		return
	}

	fmt.Println(title + "(y/n)")
	var sel string
	_, _ = fmt.Scan(&sel)
	switch strings.ToLower(sel) {
	case Y:
		fn()
	case N:
		fmt.Println("您取消了命令的执行，程序退出")
	}
}

func ConfirmOk(title string) bool {
	fmt.Println(title + "(y/n)")
	var sel string
	_, _ = fmt.Scan(&sel)
	return strings.ToLower(sel) == Y
}

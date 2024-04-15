package cmd

import (
	"fmt"
	"github.com/gophero/goal/logx"
	"os/exec"
)

func Exec(shell string, args ...string) string {
	cmd := exec.Command(shell, args...)
	logx.Default.Debugf("正在执行脚本: %s", cmd.String())
	bs, err := cmd.CombinedOutput()
	if err != nil {
		logx.Default.Errorf(fmt.Sprintf("执行脚本出错, 脚本: %s, 错误: %v", cmd.String(), err))
		return string(bs)
	}
	return string(bs)
}

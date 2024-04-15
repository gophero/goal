package cmd

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExecCmd(t *testing.T) {
	s, err := exec.Command("bash", "-c", "ps -ef | grep go").CombinedOutput()
	// s, err := exec.Command("ps -ef | grep go").CombinedOutput()
	fmt.Println(string(s), err)

	fmt.Println()
	str := Exec("bash", "-c", "ps -ef | grep go")
	fmt.Println(str)

	fmt.Println()
	str = Exec("ls", "/users")
	fmt.Println(str)

	fmt.Println()
	str = Exec("bash", "-c", "/Users/sam/tools/proxy/change_proxy.sh config_1.json > /dev/null 2>&1")
	fmt.Println(str)
}

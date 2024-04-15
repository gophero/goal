package iox

import (
	"bytes"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// Path is the global pointer which hold convenient methods for path operation.
var Path = &paths{}

type paths struct {
}

// ExecPath reports the path of current executable file, please see os.Executable for more details .
func (ps *paths) ExecPath() string {
	_path, err := os.Executable() // 获得程序路径
	if err != nil {
		panic(err)
	}
	return filepath.Dir(_path)
}

// CurrentPath reports the path of current executable file, please see runtime.Caller for more details .
func (ps *paths) CurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// PathExists reports whether the given path exists.
func (ps *paths) PathExists(path string) bool {
	return Exists(path)
}

// ProjectPath reports root path of current project. It will first execute 'go env GOMOD' command to
// get full path of 'go.mod' file of current project, if it does not exist then the method will invoke ExecPath.
func (ps *paths) ProjectPath() string {
	// default linux/mac os
	var (
		sp = "/"
		ss []string
	)

	// go env GOMOD
	// in go source code:
	// Check for use of modules by 'go env GOMOD',
	// which reports a go.mod file path if modules are enabled.

	stdout, _ := exec.Command("go", "env", "GOMOD").Output()
	p := string(bytes.TrimSpace(stdout))
	if p != "/dev/null" && p != "" {
		ss = strings.Split(p, sp)
		ss = ss[:len(ss)-1]
		return strings.Join(ss, sp)
	}
	return ps.ExecPath()
}

package iox

import (
	"github.com/gophero/goal/errorx"
	"os"
	"path/filepath"
)

var Dir = &dirs{}

type dirs struct {
}

func (d *dirs) Exists(dir string) (bool, error) {
	fi, err := os.Stat(dir)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errorx.New("file exists")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (d *dirs) AppendSep(dir string) string {
	if dir == "" {
		return dir
	}
	s := string(filepath.Separator)
	if dir[len(dir)-1:] == s {
		return dir
	} else {
		return dir + s
	}
}

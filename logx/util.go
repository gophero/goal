package logx

import (
	"errors"
	"os"
)

// PathExists 文件是否存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("file exists")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

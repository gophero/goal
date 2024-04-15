package iox

import (
	"github.com/gophero/goal/errorx"
	"os"
	"path/filepath"
	"strings"
)

// TODO 筛选过滤
func WalkDir(path string) []string {
	fi, err := os.Stat(path)
	errorx.Throw(err)

	if fi.IsDir() {
		root := Dir.AppendSep(path)
		return walkDir(root)
	} else {
		return []string{path}
	}
}

func walkDir(dir string) []string {
	dir = Dir.AppendSep(dir)
	// 读取文件
	es, err := os.ReadDir(dir)
	errorx.Throw(err)
	var files []string
	for _, entry := range es {
		if entry.IsDir() {
			files = append(files, walkDir(dir+entry.Name())...)
		} else {
			fi, _ := entry.Info()
			if strings.ToLower(filepath.Ext(entry.Name())) != ".txt" {
				continue
			}
			file := dir + fi.Name()
			files = append(files, file)
		}
	}
	return files
}

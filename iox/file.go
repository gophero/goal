package iox

import "os"

var File = &files{}

type files struct {
}

func (f *files) Exists(file string) bool {
	b := Exists(file)
	if b {
		f, _ := os.Stat(file)
		return !f.IsDir()
	}
	return b
}

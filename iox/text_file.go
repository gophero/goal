package iox

import (
	"bufio"
	"github.com/gophero/goal/errorx"
	"io"
	"os"
	"strings"
)

type TxtFile struct {
	file string
	f    *os.File
	bufw *bufio.Writer
}

func NewTxtFile(f string) *TxtFile {
	tf := &TxtFile{
		file: f,
	}
	tf.f = tf.createAndOpen(f)
	tf.bufw = bufio.NewWriter(tf.f)
	return tf
}

func (tf *TxtFile) createAndOpen(fpath string) *os.File {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errorx.Throw(err)
	return f
}

func (tf *TxtFile) WriteLine(s string) *TxtFile {
	if tf.bufw == nil {
		tf.bufw = bufio.NewWriter(tf.f)
	}
	_, err := tf.bufw.WriteString(s + "\n")
	errorx.Throw(err)
	return tf
}

func (tf *TxtFile) Flush() {
	errorx.Throw(tf.bufw.Flush())
}

func (tf *TxtFile) ReadAll() []string {
	var content []string
	var buf = bufio.NewReader(tf.f)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		content = append(content, string(line))
	}
	return content
}

// text tool

func FmtLine(line string, noSpace bool) string {
	line = strings.Trim(line, " ")
	line = strings.Trim(line, "\r")
	line = strings.Trim(line, "\n")
	if noSpace {
		line = strings.ReplaceAll(line, " ", "")
	}
	return line
}

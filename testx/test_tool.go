package testx

import (
	"strconv"
	"testing"
)

const (
	checkMark = "\u2713" // √
	ballotX   = "\u2717" // ×
)

type Logger struct {
	caseNum int
	prefix  string
	*testing.T
}

// Wrap a point of testing.T to Logger.
func Wrap(t *testing.T) *Logger {
	return &Logger{T: t}
}

// Title is a method to log current test cases(i.e. current method)'s title,
// which is a brief description for current test cases.
func (l *Logger) Title(format string, args ...any) {
	l.Logf("Test Case => "+format, args...)
}

// Case is a method to start a new test case, the format parameter will describe the case's purpose,
// given condition, and expecting results.
func (l *Logger) Case(format string, args ...any) {
	l.caseNum++
	l.prefix = "Case " + strconv.Itoa(l.caseNum) + " -> "
	l.Logf(l.prefix+format, args...)
}

// Pass is a method who invoking indicate that current condition matches the expecting result.
func (l *Logger) Pass(format string, args ...any) {
	l.Logf("\t%s "+format, prependTag(checkMark, args...)...)
}

// Fail is an opposite method to Pass, it indicate that current condition does not match the expecting result.
func (l *Logger) Fail(format string, args ...any) {
	l.Errorf("\t%s "+format, prependTag(ballotX, args...)...)
}

func (l *Logger) Quit(format string, args ...any) {
	prependTag(ballotX, args...)
	l.Fatalf("\t%s "+format, prependTag(ballotX, args...)...)
}

func prependTag(tag any, args ...any) []any {
	if args == nil {
		args = make([]any, 1)
		args[0] = tag
	} else {
		args = append([]any{tag}, args...)
	}
	return args
}

// Require is a convenient method for Pass and Fail, it requires the given bool parameter named "cond" to be true,
// if so it will invoke Pass, otherwise invoke Fail.
func (l *Logger) Require(cond bool, desc string, args ...any) {
	if cond {
		l.Pass(desc, args...)
	} else {
		l.Fail(desc, args...)
	}
}

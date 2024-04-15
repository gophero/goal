package stringx

import "strings"

func CutString(max int, s string) string {
	if len(s) > max {
		return s[:max] + "..."
	}
	return s
}

func RemSpace(s string) string {
	if s == "" {
		return s
	}
	return strings.ReplaceAll(s, " ", "")
}

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func TrimLeft(s string, cut string) string {
	return strings.TrimLeft(s, cut)
}

func Trimright(s string, cut string) string {
	return strings.TrimRight(s, cut)
}

func IsEmpty(s string) bool {
	return s == ""
}

func IsSpace(s string) bool {
	return RemSpace(s) == ""
}

func HasLen(s string) bool {
	return RemSpace(s) != ""
}

func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func ReplaceN(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

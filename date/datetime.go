package date

import (
	"math"
	"strings"
	"time"
)

type Layout string

const (
	YmdhmsEmpty = Layout("20060102150405")
	YmdEmpty    = Layout("20060102")
	YmdhmsDash  = Layout("2006-01-02 15:04:05")
	YmdDash     = Layout("2006-01-02")
	YmdhmsSlash = Layout("2006/01/02 15:04:05")
	YmdSlash    = Layout("2006/01/02")
	YmdhmsZh    = Layout("2006年01月02日 15时04分05秒")
	YmdZh       = Layout("2006年01月02日")
)

func (ly Layout) String() string {
	return string(ly)
}

func ParseDate(s string) time.Time {
	i := strings.Index(s, " ")
	if i > 0 {
		s = s[0:i]
	}

	layout := parseLayout(s)
	t, err := time.Parse(layout.String(), s)
	if err != nil {
		panic(err)
	}
	return t
}

func ParseDatetime(s string) time.Time {
	layout := parseLayout(s)
	t, err := time.Parse(layout.String(), s)
	if err != nil {
		panic(err)
	}
	return t
}

func parseLayout(s string) Layout {
	i := strings.Index(s, " ")
	incTime := i > 0
	switch {
	case strings.Index(s, "/") > 0:
		if incTime {
			return YmdhmsSlash
		}
		return YmdSlash
	case strings.Index(s, "-") > 0:
		if incTime {
			return YmdhmsDash
		}
		return YmdDash
	case strings.Index(s, "年") > 0:
		if incTime {
			return YmdhmsZh
		}
		return YmdZh
	default:
		if incTime || len(s) == 14 {
			return YmdhmsEmpty
		}
		return YmdEmpty
	}
}

func FmtDate(date time.Time) string {
	return date.Format(YmdDash.String())
}

func FmtDateTime(date time.Time) string {
	return date.Format(YmdhmsDash.String())
}

func Fmt(date time.Time, layout Layout) string {
	return date.Format(string(layout))
}

func DiffDay(src time.Time, dst time.Time) int {
	src = time.Date(src.Year(), src.Month(), src.Day(), 0, 0, 0, 0, time.Local)
	dst = time.Date(dst.Year(), dst.Month(), dst.Day(), 0, 0, 0, 0, time.Local)
	return DiffDaySec(src, dst)
}

func DiffDaySec(src time.Time, dst time.Time) int {
	return int(math.Abs(float64(int(src.Sub(dst).Seconds() / (24 * 60 * 60)))))
}

func DiffSec(src time.Time, dst time.Time) int {
	return int(math.Abs(src.Sub(dst).Seconds()))
}

func DiffMin(src time.Time, dst time.Time) int {
	return int(math.Abs(src.Sub(dst).Minutes()))
}

func DiffHour(src time.Time, dst time.Time) int {
	return int(math.Abs(src.Sub(dst).Hours()))
}

func LastDay(date time.Time) time.Time {
	return date.AddDate(0, 0, -1)
}

func NextDay(date time.Time) time.Time {
	return date.AddDate(0, 0, 1)
}

func AddDate(src time.Time, year int, month int, day int) time.Time {
	return src.AddDate(year, month, day)
}

func Add(src time.Time, d time.Duration) time.Time {
	return src.Add(d)
}

func Before(src time.Time, dst time.Time) bool {
	return src.Before(dst)
}

func After(src time.Time, dst time.Time) bool {
	return src.After(dst)
}

func Equal(src time.Time, dst time.Time) bool {
	return src.Equal(dst)
}

func EqualDate(src time.Time, dst time.Time) bool {
	if src.Year() != dst.Year() {
		return false
	}
	if src.Month() != dst.Month() {
		return false
	}
	if src.Day() != dst.Day() {
		return false
	}
	return true
}

func StartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func EndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
}

func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, -1, t.Location())
}

func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, -1, t.Location())
}

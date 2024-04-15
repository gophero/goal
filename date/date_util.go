package date

import (
	"bytes"
	"github.com/gophero/goal/conv"
	"math"
	"strconv"
	"strings"
	"time"
)

// StrTime 时间转换函数
func StrTime(datestr string, times time.Time) string {
	atime := times.Unix()
	var byTime = []int64{24 * 60 * 60, 60 * 60, 60}
	unit := strings.Split(datestr, ",")
	now := time.Now().Unix()
	ct := now - atime
	if ct < 60*5 {
		return unit[3]
	}
	//格式化成年月日
	if ct > 365*24*60*60 {
		return times.Format("2006-01-02")
	}
	//格式化成月日
	if ct > 4*24*60*60 {
		return times.Format("1-2")
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = MergeString(tempStr, unit[i]) //此处调用了一个我自己封装的字符串拼接的函数（你也可以自己实现）
		}
		break //我想要的形式是精确到最大单位，即："2天前"这种形式，如果想要"2天12小时36分钟48秒前"这种形式，把此处break去掉，然后把字符串拼接调整下即可（别问我怎么调整，这如果都不会我也是无语）
	}
	return res
}

// MergeString 拼接字符串
func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}

func ParseTime(s string) int64 {
	if strings.Index(s, ":") > 0 {
		tmp := strings.Split(s, ":")
		switch len(tmp) {
		case 1: // 异常
			return 0
		case 2:
			var s int64
			if tmp[0] != "" {
				s += conv.StrToInt64(tmp[0]) * 60
			}
			if tmp[1] != "" {
				s += conv.StrToInt64(tmp[1])
			}
			return s
		case 3:
			var s int64
			if tmp[0] != "" {
				s += conv.StrToInt64(tmp[0]) * 60 * 60
			}
			if tmp[1] != "" {
				s += conv.StrToInt64(tmp[1]) * 60
			}
			if tmp[2] != "" {
				s += conv.StrToInt64(tmp[2])
			}
			return s
		}
	}
	return conv.StrToInt64(s)
}

package date_test

import (
	"fmt"
	"github.com/gophero/goal/date"
	"reflect"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	s := "20220616"
	tm := date.ParseDate(s)
	fmt.Printf("%v\n", tm)

	tm = date.ParseDate("2022-06-16")
	fmt.Printf("%v\n", tm)

	tm = date.ParseDate("2022/06/16")
	fmt.Printf("%v\n", tm)

	tm = date.ParseDate("2022-06-16 10:10:10")
	fmt.Printf("%v\n", tm)

	tm = date.ParseDate("2022年06月16日 10:10:10")
	fmt.Printf("%v\n", tm)
}

func TestParseDatetime(t *testing.T) {
	s := "20220616000001"
	tm := date.ParseDatetime(s)
	fmt.Printf("%v\n", tm)

	tm = date.ParseDatetime("2022-06-16 00:02:01")
	fmt.Printf("%v\n", tm)

	tm = date.ParseDatetime("2022/06/16 00:02:01")
	fmt.Printf("%v\n", tm)

	tm = date.ParseDatetime("2022-06-16 10:10:10")
	fmt.Printf("%v\n", tm)

	tm = date.ParseDatetime("2022年06月16日 10时10分10秒")
	fmt.Printf("%v\n", tm)
}

func TestDiffDays(t *testing.T) {
	tm1 := date.ParseDatetime("2022-06-16 00:02:01")
	tm2 := date.ParseDatetime("2022-06-17 00:02:00")
	println(date.DiffDay(tm1, tm2))
	println(date.DiffDaySec(tm1, tm2))
	println(date.DiffSec(tm1, tm2))
	println(date.DiffMin(tm1, tm2))
	println(date.DiffHour(tm1, tm2))
}

func TestFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(date.Fmt(now, date.YmdhmsDash))
	fmt.Println(date.Fmt(now, date.YmdDash))
	fmt.Println(date.Fmt(now, date.YmdhmsEmpty))
	fmt.Println(date.Fmt(now, date.YmdEmpty))
	fmt.Println(date.Fmt(now, date.YmdhmsSlash))
	fmt.Println(date.Fmt(now, date.YmdSlash))
	fmt.Println(date.Fmt(now, date.YmdhmsZh))
	fmt.Println(date.Fmt(now, date.YmdZh))
}

func TestFmt1(t *testing.T) {
	i := 86399
	duration := time.Duration(i) * time.Second
	println(int64(duration.Hours()))
	println(int64(duration.Minutes()))
	println(int64(duration.Seconds()))

	h := i / 3600
	m := i % 3600 / 60
	s := i % 3600 % 60
	println(h)
	println(m)
	println(s)
}

func TestStartOfMonth(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// Add test cases.
		{name: "start-of-month-day", args: args{createTime(2023, 2, 1, 0, 0, 0, 0)}, want: createTime(2023, 2, 1, 0, 0, 0, 0)},
		{name: "end-of-month-day", args: args{createTime(2023, 2, 28, 23, 59, 59, 1e9-1)}, want: createTime(2023, 2, 1, 0, 0, 0, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := date.StartOfMonth(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfMonth() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}

func TestEndOfMonth(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// Add test cases.
		{name: "start-of-month-day", args: args{createTime(2023, 2, 1, 0, 0, 0, 0)}, want: createTime(2023, 2, 28, 23, 59, 59, 1e9-1)},
		{name: "end-of-month-day", args: args{createTime(2023, 2, 28, 23, 59, 59, 1e9-1)}, want: createTime(2023, 2, 28, 23, 59, 59, 1e9-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := date.EndOfMonth(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfMonth() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}

func TestStartOfYear(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// Add test cases.
		{name: "start-of-month-day", args: args{createTime(2023, 2, 1, 0, 0, 0, 0)}, want: createTime(2023, 1, 1, 0, 0, 0, 0)},
		{name: "end-of-month-day", args: args{createTime(2023, 2, 28, 23, 59, 59, 1e9-1)}, want: createTime(2023, 1, 1, 0, 0, 0, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := date.StartOfYear(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfYear() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}

func TestEndOfYear(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// Add test cases.
		{name: "start-of-month-day", args: args{createTime(2023, 2, 1, 0, 0, 0, 0)}, want: createTime(2023, 12, 31, 23, 59, 59, 1e9-1)},
		{name: "end-of-month-day", args: args{createTime(2023, 2, 28, 23, 59, 59, 1e9-1)}, want: createTime(2023, 12, 31, 23, 59, 59, 1e9-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := date.EndOfYear(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfYear() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}

func createTime(y, m, d, h, min, s, ns int) time.Time {
	return time.Date(y, time.Month(m), d, h, min, s, ns, time.Local)
}

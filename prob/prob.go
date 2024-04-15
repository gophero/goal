package prob

type Prob struct {
	Val    int    // 概率
	DisVal string // 显示值
}

func NewProb(v int, dis string) Prob {
	return Prob{Val: v, DisVal: dis}
}

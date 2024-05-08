package mathx

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FmtCommaInt(d int64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", d)
}

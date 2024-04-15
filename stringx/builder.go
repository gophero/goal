package stringx

import "strings"

type Builder struct {
	w   strings.Builder
	err error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WriteString(s string) *Builder {
	if b.err != nil {
		return b
	}
	if _, err := b.w.WriteString(s); err != nil {
		b.err = err
	}
	return b
}

func (b *Builder) Error() error {
	return b.err
}

func (b *Builder) String() string {
	return b.w.String()
}

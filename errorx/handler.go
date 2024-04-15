package errorx

// ErrorHandler 错误处理器，用来集中处理 error，而不需要冗余地多次处理 error。
// 如果 ErrorHandler 持有非 nil 的 error，则其 Do 方法将不再处理给定的函数；反之，则会将 Do 方法给定的函数参数的返回 error 赋给
// ErrorHandler 持有，以后多次调用 Do 方法则什么都不会做。最终可以通过 HasErr 方法来判断是否持有非 nil 的 error，或者通过 Done 方法来添加
// 处理钩子。
type ErrorHandler struct {
	err error
}

func NewHandler() *ErrorHandler {
	return &ErrorHandler{}
}

// Do 如果当前持有的 error 为 nil，则执行给定返回 error 的函数，否则什么也不做
func (h *ErrorHandler) Do(f func() error) *ErrorHandler {
	if h.err == nil {
		h.err = f()
	}
	return h
}

// HasErr 如果持有非 nil error，则返回 true，否则返回 false
func (h *ErrorHandler) HasErr() bool {
	return h.err != nil
}

// Err 返回持有的 error
func (h *ErrorHandler) Err() error {
	return h.err
}

// Done 方法表示处理完成的钩子函数，执行给定的函数，其参数为当前 handler 持有的 error
func (h *ErrorHandler) Done(f func(err error)) {
	f(h.err)
}

// IsPreferredErr 如果持有的 error 是首选的则返回 true，否则返回 false
func (h *ErrorHandler) IsPreferredErr() bool {
	return IsPreferred(h.err)
}

// PreferredOr 如果持有的 error 是首选的则返回，否则返回给定 error
func (h *ErrorHandler) PreferredOr(err error) error {
	if h.IsPreferredErr() {
		return h.err
	}
	return err
}

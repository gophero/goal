package httpx

func WrapErr(r *R, err error) {
	r.wrapErr(err)
}

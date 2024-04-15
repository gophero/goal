package slicex

func Collect[T any, M comparable](ts []T, f func(t T) M) []M {
	var rs []M
	for _, t := range ts {
		rs = append(rs, f(t))
	}
	return rs
}

func Each[T any](ts []T, f func(t T)) {
	for _, t := range ts {
		f(t)
	}
}

func Eachv[T any, V any](ts []T, f func(t T) V) []V {
	var vs = make([]V, len(ts))
	for i, t := range ts {
		vs[i] = f(t)
	}
	return vs
}

func Map[K comparable, V any](vs []V, f func(v V) K) map[K]V {
	var m = make(map[K]V, len(vs))
	for _, v := range vs {
		m[f(v)] = v
	}
	return m
}

func Mapv[K comparable, V any, T any](ts []T, f func(t T) (K, V)) map[K]V {
	var m = make(map[K]V, len(ts))
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}

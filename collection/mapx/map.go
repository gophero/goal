package mapx

// M is a alias for map[string]any, which is the most common map type.
type M map[string]any

// Map is a generic map type，key must be a comparable type (can use ==, != to compare)，
// v could be any type
type Map[K comparable, V any] map[K]V

func New[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
}

func (m Map[K, V]) Put(k K, v V) Map[K, V] {
	m[k] = v
	return m
}

func (m Map[K, V]) Del(k K) V {
	v := m[k]
	delete(m, k)
	return v
}

func (m Map[K, V]) Get(k K) V {
	return m[k]
}

func (m Map[K, V]) Keys() []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) Values() []V {
	var vals []V
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}

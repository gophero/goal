package redisx

type Redis interface {
	Keys(partern string) ([]string, error)
}

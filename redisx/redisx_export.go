package redisx

var MaxRetry = maxRetry

type UniqueQueue struct {
	*uniqueQueue
}

func NewExportUniqueQueue(rc Client, key string) UniqueQueue {
	return UniqueQueue{NewUniqueQueue(rc, key).(*uniqueQueue)}
}

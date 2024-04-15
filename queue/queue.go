package queue

// ==============================
// queue api
// ==============================

type Queue interface {
	Push(v ...any) error
	Pop() (string, error)
	PopN(n int) ([]string, error)
	BPop() (string, error)
	Len() uint64
}

type BoundedQueue interface {
	Queue
	Cap() uint64
}

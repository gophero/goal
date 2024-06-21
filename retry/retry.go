package retry

import (
	"log"
	"time"
)

func Do(f func() error, times int) bool {
	var n int
	for {
		if err := f(); err == nil {
			return true
		}
		n++
		if n > times {
			log.Printf("ERROR: do failed after retry the max %d times\n", times)
			return false
		}
		log.Printf("ERROR: do failed, will retry at %d time after %d second\n", n, 1<<n)
		time.Sleep(time.Second * (1 << n))
	}
}

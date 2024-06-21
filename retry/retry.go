package retry

import (
	"log"
	"time"
)

func Do(f func() error, times int) (bool, error) {
	var n int
	var err error
	for {
		if err = f(); err == nil {
			return true, nil
		}
		n++
		if n > times {
			log.Printf("ERROR: do failed after retry the max %d times\n", times)
			return false, err
		}
		log.Printf("ERROR: do failed, will retry at %d time after %d second, error: %v\n", n, 1<<n, err)
		time.Sleep(time.Second * (1 << n))
	}
}

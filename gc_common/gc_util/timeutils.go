package gc_util

import "time"

const nanoseconds = 1000000000.0

func Now() float64 {
	return float64(time.Now().UnixNano()) / nanoseconds
}

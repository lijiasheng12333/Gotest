package routine

import (
	"runtime"
	"time"
)

func Sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	println(sum)
	time.Sleep(1 * time.Second)
}

func GetRoutine() int {
	return runtime.NumGoroutine()
}

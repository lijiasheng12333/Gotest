package main

import "runtime"

func main() {
	/**
	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		time.Sleep(1 * time.Second)
	}()
	println("NewGoroutine", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	*/
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(2)
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
}

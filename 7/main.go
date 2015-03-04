package main

import "sync/atomic"

var primeCount uint64 = 1
var lastPrime int64

func prime(check int, c chan int) {
	var tmp int
	next := make(chan int)
	found := false
	for tmp = range c {
		if tmp%check != 0 {
			if !found {
				go prime(tmp, next)
				atomic.AddUint64(&primeCount, 1)
				atomic.StoreInt64(&lastPrime, int64(tmp))
				found = true
			} else {
				next <- tmp
			}
		}
	}
}

func main() {
	c := make(chan int)
	go prime(2, c)
	for i := 3; atomic.LoadUint64(&primeCount) <= 10001; i++ {
		c <- i
	}

	println(lastPrime)
}

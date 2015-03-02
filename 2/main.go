package main

const (
	LIMIT = 4000000
)

func even(i int) bool {
	return i%2 == 0
}

func fib(a, b int, r chan int) int {
	tmp := a + b
	if tmp > LIMIT {
		close(r)
		return 0
	}
	r <- tmp

	return fib(b, tmp, r)
}

func collect(c chan int) int {
	tmp := 0
	sum := 0
	for tmp = range c {
		if even(tmp) {
			sum += tmp
		}
	}
	return sum
}

func main() {
	r := make(chan int)
	go fib(1, 2, r)

	print(collect(r) + 2)
}

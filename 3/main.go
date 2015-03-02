package main

var (
	primes = map[uint64]struct{}{
		2: struct{}{},
	}
	max = uint64(GOAL)
)

const (
	GOAL = 600851475143
)

func isPrime(i uint64) bool {
	// assume our number is prime
	prime := true
	for k, _ := range primes {
		if i%k == 0 {
			prime = false
			break
		}
	}
	return prime
}

func Primes() uint64 {

	var highest uint64
	// fill in multiples of this number
	for i := uint64(3); i < max; i++ {

		// if the number is a factor of the goal test if it is a prime
		if GOAL%i == 0 {
			if isPrime(i) {
				primes[i] = struct{}{}
				max = GOAL / i
				highest = i
			}
		}
	}

	return highest
}

func main() {
	println(Primes())
}

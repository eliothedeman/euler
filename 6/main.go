package main

func sumSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += (i * i)
	}

	return sum
}

func sumNat(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func main() {
	s := sumNat(100)
	println((s * s) - sumSquares(100))

}

package main

import "math"

func sumSquares(num float64) float64 {
	sum := 0.0
	for i := 1.0; i <= num; i++ {
		sum += math.Pow(i, 2)
	}

	return sum
}

func sumNat(num float64) float64 {
	sum := 0.0
	for i := 0.0; i <= num; i++ {
		sum += i
	}
	return sum
}

func main() {

	println(math.Pow(sumNat(100), 2) - sumSquares(100))

}

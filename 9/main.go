package main

import (
	"math"
)

func main() {
	var a, b, c int
	for a < 1000 {
		b = 0
		for b < 1000 {
			c = (a * a) + (b * b)
			if math.Sqrt(float64(c)) == math.Floor(math.Sqrt(float64(c))) {
				c = int(math.Sqrt(float64(c)))
				if a+b+c == 1000 {
					println(a * b * c)
				}
			}

			b++
		}
		a++
	}
}

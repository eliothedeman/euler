package main

import "strconv"

func isPalandrome(s string) bool {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tmp string
	highest := 0
	var j int
	for i := 0; i < 900; i++ {
		for k := 0; k < 900; k++ {
			j = (100 + i) * (100 + k)
			tmp = strconv.Itoa(j)
			if isPalandrome(tmp) {
				if j > highest {
					highest = j
				}
			}
		}
	}

	println(highest)
}

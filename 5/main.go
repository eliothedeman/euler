package main

func isDivisible(j int) bool {
	for i := 1; i <= 20; i++ {
		if j%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	count := 21
	for {
		if isDivisible(count) {
			print(count)
			return
		}
		count++
	}
}

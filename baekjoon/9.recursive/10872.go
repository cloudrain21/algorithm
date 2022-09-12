package main

import . "fmt"

func f(n int) int {
	if n == 0 {
		return 1
	}
	return n * f(n-1)
}

func main() {
	var n int
	Scan(&n)
	Println(f(n))
}

package main

import . "fmt"

func b(n int) int {
	if n <= 1 {
		return n
	}
	return b(n-1) + b(n-2)
}

func main() {
	var n int
	Scan(&n)
	Println(b(n))
}
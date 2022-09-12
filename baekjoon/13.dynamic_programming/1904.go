package main

import "fmt"

var dd [1000001]int

const c = 15746

func r(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if dd[n] > 0 {
		return dd[n]
	}

	dd[n] = (r(n-1) % c + r(n-2) % c) % c
	return dd[n]
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(r(n))
}

// (a+b) % p = (a % p + b % p) % p

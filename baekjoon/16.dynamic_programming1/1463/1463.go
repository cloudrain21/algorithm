package main

import (
	"fmt"
)

var d []int

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func r(n int) int {
	if n == 1 {
		return 0
	}

	if d[n] > 0 {
		return d[n]
	}

	r1 := 1000001
	r2 := 1000001
	r3 := 1000001

	if n % 3 == 0 {
		r1 = r(n/3)
	}

	if n % 2 == 0 {
		r2 = r(n/2)
	}

	r3 = r(n-1)

	d[n] = min(min(r1,r2),r3) + 1
	return d[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	d = make([]int, n+1)
	fmt.Println(r(n))
}

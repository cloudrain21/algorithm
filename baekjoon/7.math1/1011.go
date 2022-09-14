package main

import . "fmt"

func ca(r, s int) int {
	if r <= s {
		return 1
	}
	if r <= s * 2 {
		return 2
	}

	return ca(r-s*2, s+1) + 2
}

func main() {
	var n int
	var x, y int
	Scan(&n)
	for i := 0; i < n; i++ {
		Scan(&x, &y)
		Println(ca(y-x, 1))
	}
}

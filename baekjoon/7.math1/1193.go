package main

import . "fmt"

func main() {
	var n int
	Scan(&n)

	sum := 0
	i := 1
	for ; ;i++ {
		sum += i
		if n <= sum {
			break
		}
	}

	var u, d int

	r := n - (sum - i)
	if i % 2 == 0 {
		u, d = r, i-r+1
	} else {
		u, d = i-r+1, r
	}
	Printf("%d/%d\n", u, d)
}

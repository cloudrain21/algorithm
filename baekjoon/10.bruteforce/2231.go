package main

import . "fmt"

func main() {
	var n,i,s int
	Scan(&n)

	for i=1; i<n; i++ {
		s = 0
		x := i
		r := x % 10
		for r > 0 {
			s += r
			x /= 10
			r = x % 10
		}
		if i+s == n {
			Println(i)
			return
		}
	}
	Println(0)
}

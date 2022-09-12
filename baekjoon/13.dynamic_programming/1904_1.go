package main

import . "fmt"

func main() {
	var n int
	Scan(&n)

	a, b := 1, 1
	n -= 1
	for n > 0 {
		a, b = b, (a+b)%15746
		n--
	}
	Println(b)
}

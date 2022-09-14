package main

import . "fmt"

var a = []int {
	2, 2, 2,
	3, 3, 3,
	4, 4, 4,
	5, 5, 5,
	6, 6, 6,
	7, 7, 7, 7,
	8, 8, 8,
	9, 9, 9, 9,
}

func main() {
	var s string
	Scan(&s)

	sum := 0
	for _, c := range s {
		sum += a[c-'A'] + 1
	}

	Println(sum)
}
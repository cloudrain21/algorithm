package main

import (
	. "fmt"
	"sort"
)

func main() {
	var n int
	Scan(&n)

	var a []int
	t := n
	for t > 0 {
		a = append(a, t % 10)
		t /= 10
	}
	sort.Ints(a)
	for i:=len(a)-1; i>=0; i-- {
		Print(a[i])
	}
}

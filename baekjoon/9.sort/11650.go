package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

type CD struct {
	x int
	y int
}

func main() {
	var n int

	Scan(&n)

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	a := make([]CD, n)
	for i := 0; i < n; i++ {
		Fscan(r, &a[i].x, &a[i].y)
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].x == a[j].x {
			if a[i].y < a[j].y {
				return true
			}
			return false
		}
		return a[i].x < a[j].x
	})

	for i := 0; i < n; i++ {
		Fprintln(w, a[i].x, a[i].y)
	}
	w.Flush()
}

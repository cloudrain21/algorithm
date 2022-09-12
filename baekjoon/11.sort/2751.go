package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	var n int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscanln(r, &a[i])
	}

	sort.Ints(a)
	for i := 0; i < n; i++ {
		Fprintln(w, a[i])
	}
	w.Flush()
}

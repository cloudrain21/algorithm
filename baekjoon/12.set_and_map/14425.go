package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	var t string

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	a := map[string]struct{}{}

	fmt.Fscanln(r, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &t)
		a[t] = struct{}{}
	}

	count := 0
	for i := 0; i < m; i++ {
		fmt.Fscanln(r, &t)
		if _, ok := a[t]; ok {
			count++
		}
	}
	fmt.Fprintln(w, count)
}

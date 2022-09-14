package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, t int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	a := map[int]struct{}{}

	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &t)
		a[t] = struct{}{}
	}

	fmt.Fscan(r, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &t)
		if _, ok := a[t]; ok {
			fmt.Fprint(w, "1 ")
		} else {
			fmt.Fprint(w, "0 ")
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

var d [101]int

func r(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return 1
	}

	if d[n] > 0 {
		return d[n]
	}

	d[n] = r(n-2) + r(n-3)
	return d[n]
}

func main() {
	var t,n int
	fmt.Scan(&t)

	rr := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for i:=0; i<t; i++ {
		fmt.Fscan(rr, &n)
		fmt.Fprintln(w, r(n))
	}
	w.Flush()
}

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	var n,t int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	Scan(&n)
	a := make([]int, 10001)
	for i := 0; i < n; i++ {
		Fscanln(r, &t)
		a[t]++
	}

	for i := 1; i <= 10000; i++ {
		for j := 0; j < a[i]; j++ {
			Fprintln(w, i)
		}
	}
	w.Flush()
}
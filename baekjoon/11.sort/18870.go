package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	a := make([]int, n)
	b := make([]int, n)
	m := make(map[int]int, n)

	for i:=0; i<n; i++ {
		fmt.Fscan(r, &a[i])
		b[i] = a[i]
	}

	sort.Ints(b)
	c := 0
	for i:=0; i<n; i++ {
		if _, ok := m[b[i]]; !ok {
			m[b[i]] = c
			c++
		}
	}

	for i:=0; i<n; i++ {
		fmt.Fprint(w, m[a[i]], " ")
	}
	w.Flush()
}

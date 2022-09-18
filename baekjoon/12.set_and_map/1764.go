package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, m int
	var s string

	mm := make(map[string]struct{}, 0)

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscan(r, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s)
		mm[s] = struct{}{}
	}

	rr := make([]string, 0)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &s)
		if _, ok := mm[s]; ok {
			rr = append(rr, s)
		}
	}

	sort.Strings(rr)

	fmt.Fprintln(w, len(rr))
	for _, v := range rr {
		fmt.Fprintln(w, v)
	}
	w.Flush()
}

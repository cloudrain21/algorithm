package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

type wd struct {
	w string
	l int
}

type arr []wd

func (a arr) Len() int {
	return len(a)
}

func (a arr) Less(i, j int) bool {
	if a[i].l == a[j].l {
		return a[i].w < a[j].w
	}
	if a[i].l < a[j].l {
		return true
	}
	return false
}

func (a arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	var n int
	Scan(&n)

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	wa := make(arr, n)
	for i:=0; i<n; i++ {
		Fscan(r, &wa[i].w)
		wa[i].l = len(wa[i].w)
	}

	sort.Sort(wa)

	s := ""
	for i:=0; i<n; i++ {
		if wa[i].w != s {
			Fprintln(w, wa[i].w)
			s = wa[i].w
		}
	}
	w.Flush()
}
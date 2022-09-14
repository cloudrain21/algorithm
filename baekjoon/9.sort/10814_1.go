package main

import (
	bio "bufio"
	. "fmt"
	"os"
	"sort"
)

type pp struct {
	age int
	name string
}

func main() {
	var n int
	Scan(&n)

	r := bio.NewReader(os.Stdin)
	w := bio.NewWriter(os.Stdout)

	a := make([]pp, n)
	for i:=0; i<n; i++ {
		Fscanln(r, &a[i].age, &a[i].name)
	}

	sort.SliceStable(a, func(i,j int) bool {
		return a[i].age < a[j].age
	})

	for i:=0; i<n; i++ {
		Fprintln(w, a[i].age, a[i].name)
	}
	w.Flush()
}

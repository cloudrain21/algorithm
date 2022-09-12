package main

import (
	bio "bufio"
	. "fmt"
	"os"
	"sort"
)

type p struct {
	age int
	name string
	order int
}

type ap []p

func (a ap)Len() int {
	return len(a)
}

func (a ap) Less(i, j int) bool {
	if a[i].age == a[j].age {
		return a[i].order < a[j].order
	}
	return a[i].age < a[j].age
}

func (a ap) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}


func main() {
	var n int
	Scan(&n)

	r := bio.NewReader(os.Stdin)
	w := bio.NewWriter(os.Stdout)

	a := make(ap, n)
	for i:=0; i<n; i++ {
		Fscanln(r, &a[i].age, &a[i].name)
		a[i].order = i
	}

	sort.Sort(a)

	for i:=0; i<n; i++ {
		Fprintln(w, a[i].age, a[i].name)
	}
	w.Flush()
}

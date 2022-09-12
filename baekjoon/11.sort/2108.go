package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	var n int
	Scan(&n)

	a := make([]int, n)
	m := make([]int, 8001)
	s := 0

	r := bufio.NewReader(os.Stdin)
	mcnt := 0
	for i:=0; i<n; i++ {
		Fscanln(r, &a[i])
		t := a[i] + 4000
		m[t]++
		if m[t] > mcnt {
			mcnt = m[t]
		}
		s += a[i]
	}

	var xx []int
	for i:=0; i<8001; i++ {
		if m[i] == mcnt {
			xx = append(xx, i-4000)
		}
	}

	sort.Ints(a)
	sort.Ints(xx)

	Printf("%.0f\n", float64(s)/float64(n))
	Println(a[n/2])
	if len(xx) == 1 {
		Println(xx[0])
	} else {
		Println(xx[1])
	}
	Println(a[n-1]-a[0])
}
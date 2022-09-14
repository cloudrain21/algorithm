package main

import (
	"bufio"
	"fmt"
	"os"
)

var w1 = bufio.NewWriter(os.Stdout)
var a1 []bool
var r1 []int

func f1(s, n, cw, m int) {
	if cw == m {
		for i:=0; i<len(r1); i++ {
			fmt.Fprint(w1, r1[i], " ")
		}
		fmt.Fprintln(w1)
	}

	for i := s; i<=n; i++ {
		if !a1[i] {
			a1[i] = true
			r1 = append(r1, i)
			f1(i+1, n, cw+1, m)
			a1[i] = false
			r1 = r1[:len(r1)-1]
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	for i:=0; i<=n; i++ {
		a1 = append(a1, false)
	}

	f1(1, n, 0, m)
	w1.Flush()
}

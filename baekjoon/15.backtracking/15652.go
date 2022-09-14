package main

import (
	"bufio"
	. "fmt"
	"os"
)

var w3 = bufio.NewWriter(os.Stdout)
var r3 []int

func f3(s, n, cw, m int) {
	if cw == m {
		for i:=0; i<len(r3); i++ {
			Fprint(w3, r3[i], " ")
		}
		Fprintln(w3)
		return
	}

	for i := s; i<=n; i++ {
		r3 = append(r3, i)
		f3(i, n, cw+1, m)
		r3 = r3[:len(r3)-1]
	}
}

func main() {
	var n, m int
	Scan(&n, &m)

	f3(1, n, 0, m)
	w3.Flush()
}
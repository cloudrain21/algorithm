package main

import (
	"bufio"
	. "fmt"
	"os"
)

var w2 = bufio.NewWriter(os.Stdout)
var r2 []int

func f2(n, cw, m int) {
	if cw == m {
		for i:=0; i<len(r2); i++ {
			Fprint(w2, r2[i], " ")
		}
		Fprintln(w2)
		return
	}

	for i := 1; i<=n; i++ {
		r2 = append(r2, i)
		f2(n, cw+1, m)
		r2 = r2[:len(r2)-1]
	}
}

func main() {
	var n, m int
	Scan(&n, &m)

	f2(n, 0, m)
	w2.Flush()
}

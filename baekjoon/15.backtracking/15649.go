package main

import (
	"bufio"
	. "fmt"
	"os"
)

var m int
var r []int
var a []bool
var w = bufio.NewWriter(os.Stdout)

func f(n, wd int) {
	if wd == m {
		for i:=0; i<len(r); i++ {
			Fprint(w, r[i], " ")
		}
		Fprintln(w)
		return
	}

	for i:=1; i<=n; i++ {
		if a[i] == false {
			a[i] = true
			r = append(r, i)
			f(n, wd+1)
			a[i] = false
			r = r[:len(r)-1]
		}
	}
}

func main() {
	var n int
	Scan(&n, &m)

	for i:=0; i<=n; i++ {
		a = append(a, false)
	}

	f(n, 0)
	w.Flush()
}
package main

import (
	bio "bufio"
	. "fmt"
	"os"
)

var m [15][15]int

func x(k, n int) int {
	if k == 0 {
		return n
	}
	if m[k][n] > 0 {
		return m[k][n]
	}

	sm := 0
	for i:=1; i<=n; i++ {
		sm += x(k-1,i)
	}
	m[k][n] = sm
	return sm
}

func main() {
	var t,k,n int
	Scan(&t)

	r := bio.NewReader(os.Stdin)
	for i:=0; i<t; i++ {
		Fscan(r,&k)
		Fscan(r,&n)
		Println(x(k,n))
	}
}
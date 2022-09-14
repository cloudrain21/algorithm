package main

import (
	bio "bufio"
	. "fmt"
	"os"
)

func main() {
	var n, m, t int
	Scan(&n,&m)

	r := bio.NewReader(os.Stdin)

	a := []int{}
	for i:=0; i<n; i++ {
		Fscan(r, &t)
		a = append(a, t)
	}

	s := 0
	for i:=0; i<n; i++ {
		for j:=i+1; j<n; j++ {
			for k:=j+1; k<n; k++ {
				t = a[i] + a[j] + a[k]
				if t > s && t <= m {
					s = t
				}
			}
		}
	}
	Println(s)
}

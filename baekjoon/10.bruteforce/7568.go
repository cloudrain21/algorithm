package main

import (
	"fmt"
)

type x struct {
	h int
	w int
}

func main() {
	var a []x
	var n int

	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		t := x{}
		fmt.Scan(&t.h,&t.w)
		a = append(a, t)
	}

	for i:=0; i<n; i++ {
		cnt := 0
		for j:=0; j<n; j++ {
			if a[i].h < a[j].h && a[i].w < a[j].w {
				cnt++
			}
		}
		fmt.Print(cnt+1, " ")
	}
}

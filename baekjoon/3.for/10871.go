package main

import (
	"fmt"
)

func main() {
	var n, x, t int

	fmt.Scanf("%d %d", &n, &x)

	for i:=0; i<n; i++ {
		fmt.Scan(&t)
		if t < x {
			fmt.Print(t, " ")
		}
	}
}

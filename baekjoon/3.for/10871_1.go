package main

import (
	"fmt"
)

func main() {
	var n, t int
	fmt.Scanf("%d", &n)
	for i:=0; i<n; i++ {
		fmt.Scan(&t)
		fmt.Println(t)
	}
}

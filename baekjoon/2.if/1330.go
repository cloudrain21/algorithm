package main

import "fmt"

var m = map[int]string{
	1:  ">",
	0:  "==",
	-1: "<",
}

func com(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(m[com(a, b)])
}

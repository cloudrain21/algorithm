package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	var sum, mx float64

	for i:=0; i<t; i++ {
		var c int
		fmt.Scanf("%d", &c)

		sum += float64(c) * 100
		if float64(c) > mx {
			mx = float64(c)
		}
	}

	fmt.Println(sum/mx/float64(t))
}

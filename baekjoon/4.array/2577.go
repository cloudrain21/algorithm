package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d\n%d\n%d\n", &a, &b, &c)
	t := a * b * c

	r := [10]int{}

	for {
		x := t / 10
		r[t%10]++
		if x == 0 {
			break
		}
		t = x
	}
	for i:=0; i<10; i++ {
		fmt.Println(r[i])
	}
}

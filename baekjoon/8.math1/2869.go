package main

import . "fmt"

func main() {
	var a,b,v int
	Scanf("%d %d %d\n", &a,&b,&v)

	d := (v-a)/(a-b)
	r := 0
	if (v-a)%(a-b) == 0 {
		r = d+1
	} else {
		r = d+2
	}
	Println(r)
}

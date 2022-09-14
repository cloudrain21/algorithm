package main

import . "fmt"

func main() {
	var t,h,w,n int
	Scan(&t)
	for i:=0; i<t; i++ {
		Scanf("%d %d %d\n", &h,&w,&n)
		Printf("%d%02d\n", (n-1)%h+1,(n-1)/h+1)
	}
}

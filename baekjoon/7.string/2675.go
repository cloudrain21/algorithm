package main

import . "fmt"

func main() {
	var n, c int
	var s string

	Scan(&n)
	for i:=0; i<n; i++ {
		Scan(&c)
		Scan(&s)
		for j:=0; j<len(s); j++ {
			for k:=0; k<c; k++ {
				Print(s[j:j+1])
			}
		}
		Print(" ")
	}
}

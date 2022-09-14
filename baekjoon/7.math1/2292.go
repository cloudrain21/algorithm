package main

import . "fmt"

func main() {
	var n int
	Scan(&n)

	e := 1
	for i:=0; ; i++ {
		e += 6 * i
		if n <= e {
			Println(i+1)
			break
		}
	}
}

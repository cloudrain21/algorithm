package main

import (
	. "fmt"
	. "strconv"
	. "strings"
)

func main() {
	var n int
	Scan(&n)

	i, c := 666, 0
	for c < n {
		s := Itoa(i)
		if Contains(s, "666") {
			c++
		}
		i++
	}
	Println(i-1)
}

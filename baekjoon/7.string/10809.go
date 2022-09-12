package main

import (
	. "fmt"
	. "strings"
)

func main() {
	var s string
	Scan(&s)

	b := int('a')
	for i:=b; i<b+26; i++ {
		Print(Index(s, string(i)), " ")
	}
}

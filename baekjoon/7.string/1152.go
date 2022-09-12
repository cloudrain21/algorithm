package main

import (
	bio "bufio"
	. "fmt"
	"os"
	. "strings"
)

func main() {
	var s string

	r := bio.NewReader(os.Stdin)
	b, _, _ := r.ReadLine()
	s = string(b)

	s = Trim(s, " ")
	ss := Split(s, " ")

	Print(len(ss))
}
package main

import . "fmt"

func main() {
	var a, b string
	Scanf("%s %s\n", &a, &b)

	ra := []rune(a)
	rb := []rune(b)

	ra[0], ra[2] = ra[2], ra[0]
	rb[0], rb[2] = rb[2], rb[0]

	if string(ra) > string(rb) {
		Print(string(ra))
	} else {
		Print(string(rb))
	}
}
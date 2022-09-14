package main

import (
	"fmt"
)

func reverse(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}

func ff(a, b, o int) (ro int, r string) {
	t := a + b + o
	ro = t / 10
	r = string((t % 10) + '0')
	return ro, r
}

func main() {
	var n1, n2 string
	var r, rr string

	fmt.Scanf("%s %s\n", &n1, &n2)

	var i, j int
	o := 0
	for i, j = len(n1)-1, len(n2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		a := int(n1[i] - '0')
		b := int(n2[j] - '0')
		o, rr = ff(a, b, o)
		r += rr
	}

	if i >= 0 {
		for ; i >= 0; i-- {
			a := int(n1[i] - '0')
			o, rr = ff(a, 0, o)
			r += rr
		}
	}

	if j >= 0 {
		for ; j >= 0; j-- {
			a := int(n2[j] - '0')
			o, rr = ff(a, 0, o)
			r += rr
		}
	}
	if o > 0 {
		r += string(o + '0')
	}
	fmt.Println(reverse(r))
}

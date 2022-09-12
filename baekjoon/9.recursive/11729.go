package main

import (
	"bytes"
	. "fmt"
)

var cnt = 0

func m(n, f, t int, b *bytes.Buffer) {
	if n == 1 {
		cnt++
		b.WriteString(Sprintf("%d %d\n", f, t))
		return
	}

	v := 6 - f - t
	m(n-1, f, v, b)
	cnt++
	b.WriteString(Sprintf("%d %d\n", f, t))
	m(n-1, v, t, b)
}

func main() {
	var n int
	Scan(&n)

	var b bytes.Buffer
	m(n, 1, 3, &b)
	Println(cnt)
	Print(b.String())
}

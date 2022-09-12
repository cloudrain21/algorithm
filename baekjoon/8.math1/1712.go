// timeout
package main

import "fmt"

var a,b,c int64
func f(s, e int64) int64 {
	if s >= e {
		return s
	}

	h := (s + e)/2

	ct,ip := calc(h)

	if ct >= ip {
		return f(h+1, e)
	} else {
		return f(s, h-1)
	}
}

func calc(i int64) (cost int64, input int64) {
	return a + i * b, i * c
}

func main() {
	fmt.Scanf("%d %d %d\n",&a,&b,&c)

	if c <= b {
		fmt.Println(-1)
		return
	}

	i := int64(1)
	ct,ip := calc(i)
	for ct >= ip {
		i *= 2
		ct,ip = calc(i)
	}

	fmt.Println(f(i/2,i))
}

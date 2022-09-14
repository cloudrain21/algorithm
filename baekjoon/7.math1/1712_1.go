package main

import . "fmt"

func main() {
	var a,b,c int
	Scanf("%d %d %d\n",&a,&b,&c)

	if c <= b {
		Println(-1)
		return
	}

	i := 1
	for ; i<=2100000000; i++ {
		ct, ip := a + i * b, i * c
		if ip > ct {
			break
		}
	}

	Println(i)
}

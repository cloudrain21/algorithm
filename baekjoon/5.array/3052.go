// map 을 이용하면 더 짧아질 수 있다. 필요할 때만 map[x] = true 로 하면 되니까.
package main

import "fmt"

func main() {
	var c int
	a := [42]int{}

	for i:=0; i<10; i++ {
		fmt.Scanf("%d\n", &c)
		a[c%42]++
	}
	c = 0
	for i:=0; i<42; i++ {
		if a[i] > 0 {
			c++
		}
	}
	fmt.Println(c)
}

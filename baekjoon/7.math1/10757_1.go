package main

import (
	. "fmt"
	. "math/big"
)

func main() {
	var a, b Int
	Scanf("%s %s\n", &a, &b)
	r := &Int{}
	Print(r.Add(&a,&b).String())
}

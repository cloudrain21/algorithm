package main

import (
	bio "bufio"
	. "fmt"
	"math/big"
	"os"
)

func main() {
	r := bio.NewReader(os.Stdin)
	var a, b big.Int
	Fscan(r, &a, &b)
	Println(a.Add(&a,&b))
}

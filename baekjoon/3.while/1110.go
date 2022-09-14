package main

import "fmt"

var count int

func main() {
	var n int
	fmt.Scanf("%d", &n)
	original := n

	newNum := getNewNum(n)
	for newNum != original {
		n = newNum
		newNum = getNewNum(n)
	}
	fmt.Println(count)
}

func getNewNum(n int) int {
	n1 := n % 10
	n2 := (n / 10) + (n % 10)
	count++
	return n1*10 + n2%10
}

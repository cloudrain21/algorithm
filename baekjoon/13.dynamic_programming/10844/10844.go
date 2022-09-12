package main

import "fmt"

func r(n int) int {
	if n == 1 {
		return 9
	}

	if n == 2 {
		return 17
	}

}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(r(n)/1000000000)
}

// 1  0  1
//    2  1
//       3
// 2  1  0
//       2
//    3
// 3  2
//    4
// 4  3
//    5
// 5  4
//    6
// 6  5
//    7
// 7  6
//    8
// 8  7
//    9
// 9  8

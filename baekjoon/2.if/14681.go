package main

import "fmt"

func getQuadrant(x, y int) int {
	if x > 0 && y > 0 {
		return 1
	} else if x < 0 && y > 0 {
		return 2
	} else if x < 0 && y < 0 {
		return 3
	} else {
		return 4
	}
}

func main() {
	var x, y int
	for {
		fmt.Scanf("%d\n%d", &x, &y)
		if x == 0 || x > 1000 || x < -1000 {
			continue
		}
		if y == 0 || y > 1000 || y < -1000 {
			continue
		}
		break
	}

	fmt.Println(getQuadrant(x, y))
}

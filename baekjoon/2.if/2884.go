package main

import "fmt"

func calc(h, m, ahead int) (int, int) {
	if m >= ahead {
		return h, m - ahead
	} else {
		rm := (60 + m) - ahead
		if h > 0 {
			return h - 1, rm
		} else {
			return 23, rm
		}
	}
}

func main() {
	var h, m int
	for {
		fmt.Scanf("%d %d", &h, &m)
		if h < 0 || h > 24 {
			continue
		}
		if m < 0 || m > 59 {
			continue
		}
		break
	}

	rh, rm := calc(h, m, 45)
	fmt.Println(rh, rm)
}

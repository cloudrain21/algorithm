package main

import (
	. "fmt"
	"regexp"
)

func main() {
	var s string
	Scan(&s)
	r := regexp.MustCompile("c[=-]|dz=|d-|lj|nj|[sz]=")
	rs := r.ReplaceAllString(s, ".")
	Println(len(rs))
}

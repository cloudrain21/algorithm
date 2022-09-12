package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func saveFile(name string, total int) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	rand.Seed(time.Now().UnixNano())
	for i:=0; i<total; i++ {
		n := rand.Intn(1000000)
		if n % 2 == 0 {
			n = -n
		}
		_, err := f.WriteString(strconv.Itoa(n))
		if err != nil {
			return err
		}

		if i < (total - 1) {
			_, err = f.WriteString(" ")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func readStringFromFile(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf = make([]byte, 10000000)
	n, err := f.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func removeFile(name string) {
	err := os.Remove(name)
	if err != nil {
		panic(err)
	}
}

func main() {
	var total int
	//var fileName = "./text.txt"
	fmt.Scanf("%d", &total)

	r := bufio.NewReader(os.Stdin)
	l, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	ls := strings.Trim(string(l), " ")

	//err := saveFile(fileName, total)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//ls, err := readStringFromFile(fileName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//removeFile(fileName)

	max := -1000001
	min := 1000001

	ss := strings.Split(ls, " ")
	for i := 0; i < len(ss); i++ {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			log.Fatal(err)
		}

		if n >= max {
			max = n
		}

		if n <= min {
			min = n
		}
	}

	fmt.Println(min, max)
}

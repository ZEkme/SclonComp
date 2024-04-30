package main

import (
	"fmt"
	"strconv"
)

func foo(n int) string {
	switch n % 10 {
	case 1:
		if n%100 == 11 {
			return strconv.Itoa(n) + " компьютеров"
		}
		return strconv.Itoa(n) + " компьютер"
	case 2, 3, 4:
		if n%100 == 12 || n%100 == 13 || n%100 == 14 {
			return strconv.Itoa(n) + " компьютеров"
		}
		return strconv.Itoa(n) + " компьютера"
	default:
		return strconv.Itoa(n) + " компьютеров"
	}
}

func main() {
	fmt.Println(foo(1))
	fmt.Printf("%T\n", foo(2))
}

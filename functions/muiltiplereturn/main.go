package main

import (
	"fmt"
)

func main() {
	a, b := getMessage()
	fmt.Println(a, b)

	x, y := split(100)
	fmt.Println(x, y)
}

func getMessage() (string, string) {
	return "hello", "hi"
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

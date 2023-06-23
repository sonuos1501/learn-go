package main

import (
	"fmt"
)

func fibonaci() func() int {
	return func () int  {
		return 1
	}
}

func main() {

	f := fibonaci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

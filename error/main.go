package main

import "fmt"

func main() {
	f(b)
}

func f(a func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd in f", r)
		}
	}()
	fmt.Println("Calling func a.")
	a()
	fmt.Println("Return normally from a.")
}

func b() {
	defer fmt.Println("defer")
	panic(fmt.Sprintln("panic"))
	// return "return"
}

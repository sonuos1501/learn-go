package main

import "fmt"

func main() {
	message := greetMe("Phạm Thế Sơn")

	fmt.Println(message)
}

func greetMe(name string) string {
	return "Hello, " + name + "!"
}

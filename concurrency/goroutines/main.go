package main

import "fmt"

func main() {
	ch := make(chan string)

	go push("SON", ch)
	go push("Thuy", ch)
	go push("Trang", ch)

	fmt.Println(<-ch, <-ch, <-ch)
}

func push(name string, ch chan string) {
	msg := "Hey, " + name
	ch <- msg
}
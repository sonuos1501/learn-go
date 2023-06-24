package main

import "fmt"

func main() {
	coffee := bridge("Coffee")
	bread := bridge("Bread")

	serve := fanIn(coffee, bread)

	for i := 0; i < 5; i++ {
		fmt.Println(<-serve)
	}
}


// Fan-in pattents điều hướng các channels
func fanIn(chan1, chan2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case <-chan1:
				c <- <-chan1
			case <-chan2:
				c <- <-chan2
			}
		}
	}()
	return c
}


// Return pattents
func bridge(msg string) chan string {
	result := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			result <- fmt.Sprintf("%s %d",msg, i)
		}
	}()

	return result
}
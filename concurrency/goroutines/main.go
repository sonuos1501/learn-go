package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	go printNumber(&wg)
	go printChar(&wg)

	wg.Wait()

	fmt.Println("Finish")
}

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		fmt.Printf("%d", i)
	}
}

func printChar(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c", i)
	}
}

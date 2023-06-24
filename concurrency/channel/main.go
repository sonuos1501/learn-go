package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	chanPrintNum := make(chan int)
	chanPrintChar := make(chan string)

	wg.Add(2)

	go printChar(&wg, &chanPrintChar)
	go printNumber(&wg, &chanPrintNum)

	fmt.Println("Result number", <-chanPrintNum)
	fmt.Println("Result charactor", <-chanPrintChar)

	wg.Wait()
	fmt.Println("Finished")
}

func printNumber(wg *sync.WaitGroup, numChan *chan int) {
	defer wg.Done()

	result := 0

	for i := 0; i < 100; i++ {
		result += i
	}

	*numChan <- result
}

func printChar(wg *sync.WaitGroup, strChan *chan string) {
	defer wg.Done()

	result := ""
	
	for i := 'A'; i < 'A' + 26; i++ {
		result += string(i)
	}

	*strChan <- result
}
package main

import "fmt"

func main() {
	fistChan := firstInput(100, 2, 3, 4, 5, 6)
	secondChan := secondInput(fistChan)

	for item := range secondChan {
		fmt.Println("recive", item)
	}
}

func firstInput(nums ...int) chan int {
	result := make(chan int)

	go func() {
		for _, item := range nums {
			result <- item
		}

		close(result)
	}()

	return result
}

func secondInput(fromFirst chan int) chan int {
	result := make(chan int)

	go func() {
		for item := range fromFirst {
			result <- item * item
		}
		close(result)
	}()

	return result
}

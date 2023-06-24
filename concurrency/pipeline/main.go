package main

import "fmt"

func main() {
	fistChan := firstInput(1, 2, 3, 4, 5, 6)
	secondChan := secondInput(fistChan)

	for item := range secondChan {
		fmt.Println("recive", item)
	}
}

func firstInput(nums ...int) chan int {
	result := make(chan int)

	go func() {
		for i := 0; i < len(nums); i++ {
			result <- nums[i]
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

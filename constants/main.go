package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	const Phi = 1.618
	fmt.Println(Phi)

	const Size int64 = 1024
	fmt.Println(Size)

	const x, y = 1, 2
	fmt.Println(x)
	fmt.Println(y)

	const (
		Pi = 3.14
		E  = 2.17
	)

	fmt.Println(Pi)
	fmt.Println(E)

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

}

package main

import "fmt"

// Khai báo tổng
var (
	x          string
	n          = 20
	y, yy, yyy = 1, 2.3, "HI"
)

func main() {
	// Kiểu khai báo biến hay dùng
	t := 100
	fmt.Printf("Kiểu khai báo biến hay dùng: %d\n", t)

	hh, hhh := 5, "Nguyễn Thu Thuỷ"
	fmt.Printf("Kiểu khai báo biến hay dùng, hh: %d, hhh: %s \n", hh, hhh)

	// Kiểu 1
	var a string
	a = "Kieu 1"
	fmt.Println(a)

	// Kiểu 2
	var b = "Kieu 2"
	fmt.Println(b)

	// Kiểu 3
	var c string = "Kiểu 3"
	fmt.Println(c)

	// Kiểu 4
	var d, e int
	d = 10
	e = 15
	fmt.Printf("Kiểu 4, d: %d, e: %d\n", d, e)

	// Kiểu 5
	var f, g int = 1, 2
	fmt.Printf("Kiểu 5, f: %d, g: %d\n", f, g)

	// KIểu 6
	var h, i = 3, "Hello"
	fmt.Printf("Kiểu 6, h: %d, i: %s\n", h, i)

}

package main

import "fmt"

func main() {
	x := 1
	y := 2

	result := x + y
	fmt.Println("suma:", result)

	result = x - y
	fmt.Println("resta:", result)

	result = x * y
	fmt.Println("mult:", result)

	result = x / y
	fmt.Println("div:", result)

	result = x % y
	fmt.Println("modulo:", result)

	x++
	fmt.Println("incrementar:", x)

	x--
	fmt.Println("decrementar:", x)
}

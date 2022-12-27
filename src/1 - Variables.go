package main

import "fmt"

func main() {
	// Declaracion de una constante
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println("base: ", base)
	fmt.Println("altura: ", altura)
	fmt.Println("area: ", area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("int: ", a)
	fmt.Println("float: ", b)
	fmt.Println("string: ", c)
	fmt.Println("bool: ", d)
}

package main

import "fmt"

func main() {
	// Defer hace que se ejecute una linea antes de morir ese codigo
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 8 {
			fmt.Println("Es 8")
			break
		}
	}
}

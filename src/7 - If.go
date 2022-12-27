package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	val1 := 1
	val2 := 2

	if val1 == 1 && val2 != 2 {
		fmt.Println("Las dos condiciones cumplen")
	} else {
		fmt.Println("Alguna de las condiciones no cumple")
	}

	if val1 == 1 || val2 != 2 {
		fmt.Println("Alguna de las condiciones cumple")
	} else {
		fmt.Println("Ninguna de las condiciones cumple")
	}

	// Uso tipico del if
	value, err := strconv.Atoi("14")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}

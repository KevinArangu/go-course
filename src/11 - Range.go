package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"Hola", "ke", "ase"}

	// For donde se usan los dos valores
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// For donde se usa solo el indice
	// for i := range slice {
	// 	fmt.Println(i)
	// }

	// For donde se usa solo el valor
	// for _, valor := range slice {
	// 	fmt.Println(valor)
	// }

	fmt.Println(isPalindromo("Ama"))
}

func isPalindromo(text string) bool {
	var textLower = strings.ToLower(text)
	var textReversed string

	for i := len(textLower) - 1; i >= 0; i-- {
		textReversed += string(textLower[i])
	}

	return textLower == textReversed
}

package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["Jose"] = 10
	myMap["Luis"] = 20
	fmt.Println(myMap)

	// Recorrer un map
	for i, v := range myMap {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value := myMap["Luis"]
	fmt.Println(value)

	// Si envias una llave que no existe, devuelve el Zero Value del tipo
	zeroValue := myMap["Joseph"]
	fmt.Println(zeroValue)

	// Para saber cuando una llave existe en el Map (Y asi saber si tiene un valor fallback o Zero Value)
	value2, ok := myMap["Juan"]
	fmt.Println(value2, ok)
}

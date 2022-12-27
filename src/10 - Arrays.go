package main

import "fmt"

func main() {
	// Array: Se le indica la cantidad de valores que van a existir
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	// len() Cuantos elementos hay en un array
	fmt.Println(len(array))

	// cap() Cual es la capacidad maxima de un array
	fmt.Println(cap(array))

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Metodos en el slice
	// NOTA: [ EL1 : EL2 ] El valor del EL1 es inclusivo, mientras que el valor de EL2 no lo es.
	// Por ende, al decir que va de [ 2 : 4 ], empieza del valor 2 y termina en el 3

	fmt.Println(slice[0])   // Imprimir el elemento 0
	fmt.Println(slice[:3])  // Imprimir hasta el elemento 3
	fmt.Println(slice[2:4]) // Imprimir desde el elemento 2 hasta el 4
	fmt.Println(slice[4:])  // Imprimir desde el elemento 4 en adelante
}

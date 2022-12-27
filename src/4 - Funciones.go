package main

import "fmt"

func print(value, other string) {
	fmt.Println(value, other)
}

func returnValue(value int) int {
	return value * 2
}

func doubleReturn(value int) (a, b int) {
	return value, value * 2
}

func main() {
	print("Hola", "mundo")

	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := doubleReturn(2)
	fmt.Println(value1, value2)

	other, _ := doubleReturn(8)
	fmt.Println(other)
}

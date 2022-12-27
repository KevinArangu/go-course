package main

import "fmt"

func main() {
	val1 := 1
	val2 := 2

	condition := val1 == val2
	fmt.Println(condition)

	condition = val1 != val2
	fmt.Println(condition)

	condition = val1 < val2
	fmt.Println(condition)

	condition = val1 > val2
	fmt.Println(condition)

	condition = val1 <= val2
	fmt.Println(condition)

	condition = val1 >= val2
	fmt.Println(condition)

	operator := true && false
	fmt.Println(operator)

	operator = true || false
	fmt.Println(operator)

	operator = !false
	fmt.Println(operator)

}

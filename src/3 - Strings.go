package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "World"
	fmt.Println(helloMessage, worldMessage)

	academy := "Platzi"
	numberOfCourses := 500
	fmt.Printf("%s tiene mas de %d cursos\n", academy, numberOfCourses)
	fmt.Printf("%v tiene mas de %v cursos\n", academy, numberOfCourses)

	message := fmt.Sprintf("%s tiene mas de %d cursos", academy, numberOfCourses)
	fmt.Println(message)

	fmt.Printf("academy: %T\n", academy)
	fmt.Printf("numberOfCourses: %T\n", numberOfCourses)
}

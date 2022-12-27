package myPackage

import "fmt"

// Primera letra Mayus: Publico
// Primera letra Minus: Privado

// Car car struct public
type Car struct {
	Brand string
	Year  int
}

func SayHello(name string) {
	fmt.Println(name)
}

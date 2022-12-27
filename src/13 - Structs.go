package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	// Instanciar Struct
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2022
	fmt.Println(otherCar)
}

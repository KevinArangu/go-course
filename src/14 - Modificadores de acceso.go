package main

import (
	"fmt"

	pk "mypackage"
)

func main() {
	var myCar pk.Car
	myCar.Brand = "Ferrari"
	myCar.Year = 2022

	fmt.Println(myCar)
}

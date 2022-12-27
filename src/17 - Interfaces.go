package main

import "fmt"

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.altura * r.base
}

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type figura2D interface {
	area() float64
}

func calcular(f figura2D) float64 {
	return f.area()
}

func main() {
	c := cuadrado{base: 4}
	r := rectangulo{base: 2, altura: 4}

	fmt.Println(calcular(c))
	fmt.Println(calcular(r))

	// Lista de interfaces
	myInterface := []interface{}{"Hola", 12, true}
	fmt.Println(myInterface...)
}

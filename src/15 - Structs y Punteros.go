package main

import "fmt"

type pc struct {
	brand     string
	processor string
	ram       int
	disk      int
}

func (pc pc) getBrand() string {
	return pc.brand
}

func (pc pc) getProcessor() string {
	return pc.processor
}

func (pc *pc) setBrand(newValue string) string {
	pc.brand = newValue
	return pc.brand
}

func main() {
	a := 50
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	myPc := pc{brand: "MSI", processor: "Intel i5 12gen", ram: 16, disk: 256}
	myPc.setBrand("HP")
	fmt.Println(myPc)
}

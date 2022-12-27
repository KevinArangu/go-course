package main

import "fmt"

type pc struct {
	brand string
	ram   int
	disk  int
}

func (pc pc) String() string {
	return fmt.Sprintf("Brand: %s, RAM: %d GB, Disk: %d GB", pc.brand, pc.ram, pc.disk)
}

func main() {
	myPc := pc{brand: "Lenovo", ram: 8, disk: 128}
	fmt.Println(myPc)
}

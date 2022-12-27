package main

import (
	"fmt"
	"sync"
)

func say(text string) {
	fmt.Println(text)
}

func sayConcurrency(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	// Funciona pero no es correcto
	// say("hello")
	// go say("World")
	// time.Sleep(time.Second * 1)

	// Funciona y es correcto
	var wg sync.WaitGroup
	say("hello")
	wg.Add(1)
	go sayConcurrency("World", &wg)
	wg.Wait()

	// Go routine con funciones anonimas
	// go func(text string) {
	// 	fmt.Println(text)
	// }("texto")
}

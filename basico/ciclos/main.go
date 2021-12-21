package main

import "fmt"

func main() {
	//	For condicional
	for i := 0; i < 10; i++ {
		//fmt.Printf("Desde ciclo for condicional %d\n", i)
	}

	//for while
	counter := 0
	for counter < 10 {
		fmt.Printf("Desde ciclo for while %d \n", counter)
		counter++
	}
	//For forever

	counterForever := 0
	for {
		fmt.Printf("Desde ciclo forEver %d", counterForever)
		counterForever++
	}

}

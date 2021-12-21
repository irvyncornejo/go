package main

import "fmt"

func main() {
	//modulo := 5 % 2 la condición puedo estar en una línea adelante o en la misma línea del switch
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//sin condición
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}
}

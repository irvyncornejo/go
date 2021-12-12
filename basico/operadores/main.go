package main

import "fmt"

func main() {
	//operadores
	x := 10
	y := 20
	//suma
	result := x + y
	fmt.Println("Suma: ", result)

	//resta
	result = x - y
	fmt.Println("Resta: ", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	//División
	result = y / x
	fmt.Println("División: ", result)

	//Modulo
	result = x % y
	fmt.Println("Módulo: ", result)

	//incremental
	x++
	fmt.Println("Incremental", x)

	//decremental
	x--
	fmt.Println("Decremental", x)

}

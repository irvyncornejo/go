package main

import "fmt"

func main() {
	//defer lineas que se ejecutan antes de terminar la ejecución de la función
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("Break")
			break
		}

	}

}

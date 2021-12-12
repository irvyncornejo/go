package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

//Buena practica de parametros a, b int, c string
func tripeArgumrnt(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 4
}

func doubleValues(a int) (c, d int) {
	return a, a * 4
}

func main() {
	normalFunction("Hola mundo")
	tripeArgumrnt(1, 2, "Tres")
	var value int = returnValue(3)
	fmt.Println(value)

	value1, value2 := doubleValues(3)
	fmt.Println("Value1 y value2:", value1, value2)

	//Si solo necesitamos un valor de los dos que regresa la función podemos hacer lo siguiente
	valuea, _ := doubleValues(2)
	fmt.Println("Value a:", valuea)
}

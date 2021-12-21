package main

import "fmt"

func main() {
	//Los Array son inmutables y los slice son mutables
	var array [4]int8
	array[0] = 3
	array[1] = 4
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int8{0, 2, 4, 5, 7, 9}
	fmt.Println(slice, len(slice), cap(slice))

	//Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 10)
	fmt.Println(slice)

	//Append nueva lista
	nuevaSlice := []int8{11, 23, 45, 6}
	slice = append(slice, nuevaSlice...)
	fmt.Println(slice)
}

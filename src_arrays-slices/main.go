package main

import "fmt"

func main() {

	//Array
	fmt.Println("array")
	var arreglo [4]int
	arreglo[0] = 1
	arreglo[1] = 66
	//len = cuantos elementos hay , cap = capacidad maxima del array
	fmt.Println(arreglo, len(arreglo), cap(arreglo))

	//Slice
	fmt.Println("slice")
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos slice
	fmt.Println(slice[0])
	//Imprime los 3 primeros
	fmt.Println(slice[:3])
	//imprime del 2 al 4
	fmt.Println(slice[2:4])
	//Imprime del 4 en adelante
	fmt.Println(slice[4:])

	//Append
	fmt.Println("\nappend")
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append nueva lista . unir
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)

}

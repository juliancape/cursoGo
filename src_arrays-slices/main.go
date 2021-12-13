package main

//Julian Cardenas Peñuela
//12-12-2021
import (
	"fmt"
	"strings"
)

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
	fmt.Println(slice)
	slice = append(slice, 111)
	fmt.Println(slice)

	//-----------------------------------------
	//Recorrido  de slice con range

	fmt.Println("\n--Recorrido de slice con range--")
	slice2 := []string{"hola", "que", "hace"}

	for i, valor := range slice2 {
		fmt.Println(i, valor)
	}

	//Ejercio 1
	fmt.Println("\nEjercicio Palindromo")
	slicePalindromo := []string{"carro", "ojO", "avion", "Oso", "ama"}

	for i := range slicePalindromo {
		convMinuscula := strings.ToLower(slicePalindromo[i])
		esPalindromo(convMinuscula)

	}

}

//Funcion para saber si una palabra es palindromo o no
func esPalindromo(text string) {

	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if textReverse == text {
		fmt.Println(text, "Es palindromo")
	} else {
		fmt.Println(text, "No es palindromo")
	}
}

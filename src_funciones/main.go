package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64 //Decimales
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//Numeros Enteros
	//int8 = 8bits = -128 a 127
	//int 16
	//int 32
	//int 64 = 64bits = -2^63 a 2^63 -1

	//Numeros enteros positivos
	//uint = Depende del sistema operativo
	//uint32 = 32bits = 0 a 2^32 -1
	//uint64 = 64bits = 0 a 2^64 -1

	//Numeros decimales
	//float32 = 32bits
	//float64= 64bits

	//Numeros complejos
	//Complex64 = Real e imaginario float32
	//Complex128 = Real e imaginario float64
	//Ejemplo: c:= 10-8i

	//Declaracion variables
	helloMessage := "Hello"
	worldMessage := "World"
	fmt.Println(helloMessage, worldMessage)
	nombre := "Platzy"
	cursos := 500
	//Printf
	// %s string     %d int
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	//%v no saber que tipo de datos estan ahi
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Saber tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}

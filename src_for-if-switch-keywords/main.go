package main

import "fmt"

func main() {
	//ciclo for imprime de 0 a 9
	fmt.Println("For de-hasta")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//For while
	fmt.Println("for while")
	counter := 15
	for counter >= 0 {
		fmt.Println(counter)
		counter--
	}

	//For forever
	fmt.Println("For forever")
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		//condicional pa que no se haga infinito el ciclo for forever
		if counterForever == 20 {
			break
		}
	}

	valor1 := 1
	valor2 := 2
	//condicional and
	fmt.Println("Condicional and")
	if valor1 < valor2 && valor2 == 2 {
		fmt.Println("Valor 1 es menor a valor 2 y valor 2 es igual a 2")
	} else {
		fmt.Println("Valor 1 es mayor a valor 2 y valor 2 es igual a 2")
	}
	//Condicional or
	fmt.Println("Condicional or")
	if valor1 == 1 || valor2 == 3 {
		fmt.Println("Una de las dos condiciones se cumplio")
	}

	//Condiciones anidadas con switch
	fmt.Println("switch")
	//Si modulo es igual a 0 significa que el numero es par

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//switch sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")

	}

	//keywords defer, break y continue

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("Es dos")
			continue
			//fmt.Println("Esto no se va a mostrar")
		}

		if i == 8 {
			break
		}
	}

	//defer ejecuta como su fuera la ultima linea de codigo que se lee
	defer fmt.Println("hola")
	fmt.Println("mundo")

}

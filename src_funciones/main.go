package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (b, c int) {
	return a, (a * 5) - a
}

func main() {
	normalFunction("Hola mundo")
	fmt.Println(returnValue(4))
	value1, value2 := doubleReturn(3)
	fmt.Println("Value1", value1, "Value2", value2)

}

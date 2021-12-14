package main

//Julian Cardenas Peñuela
//13-12-2021
import "fmt"

//rectangulo
type rectangulo struct {
	base   float64
	altura float64
}

//cuadrado
type cuadrado struct {
	base float64
}

type figuras2D interface {
	area() float64
}

//funcion calcular area rectangulo
func (miRectangulo rectangulo) area() float64 {
	return miRectangulo.base * miRectangulo.altura
}

//funcion calcular area cuadrado
func (miCuadrado cuadrado) area() float64 {
	return miCuadrado.base * miCuadrado.base
}

func calculate(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func main() {

	miRectangulo := rectangulo{base: 22.2, altura: 44.1}
	miCuadrado := cuadrado{base: 55.3}

	calculate(miCuadrado)
	calculate(miRectangulo)

	//--------------------
	//Lista de interfaces
	fmt.Println("Lista interfaces")
	myInterface := []interface{}{"Hola", 12, 33.2, true}
	fmt.Println(myInterface...)

}

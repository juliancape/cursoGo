package main

import (
	"fmt"

	pk "example.com/curso_go_platzy/src_acceso-func-struc-punteros/mypackage"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

//funcion con struc
func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")

}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2

	fmt.Println("RAM duplicada")
}

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	//imprimir desde otro modulo
	pk.Message("Hola como van")

	//------------------
	fmt.Println("\nPUNTEROS")
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 150, brand: "ASUS"}

	fmt.Println(myPC)
	myPC.ping()
	myPC.duplicateRAM()

	fmt.Println(myPC)

}

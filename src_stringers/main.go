package main

//Julian Cardenas Peñuela
//13-12-2021
import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	//Personalizar output
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s.", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, brand: "ASUS", disk: 100}
	fmt.Println(myPC)
}

package mypackage

//Julian Cardenas Peñuela
//13-12-2021
import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// dataPeronsaPrivate con acceso privado
type dataPersonaPrivate struct {
	nombre string
	id     int
}

//Imprimir mensaje
func Message(text string) {
	fmt.Println(text)
}

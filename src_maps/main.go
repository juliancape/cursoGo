package main

//Julian cardenas peñuela
//12-12-2021
import "fmt"

func main() {
	m := make(map[string]int)

	m["Julian"] = 18
	m["Juan"] = 22
	fmt.Println(m)

	//recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["Julian"]
	fmt.Println("Valor:", value, ok)

}

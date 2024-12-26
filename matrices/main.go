package main

import "fmt"

func main() {
	// cuando no sabemos con cuantos datos se va a inicializar un array le ponemos '...' en lugar del tamaño
	a := [...]int{10, 20, 30, 40, 50}
	a[0] = 100
	a[1] = 200
	//fmt.Println(a)
	//fmt.Println(a[1])
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for index, value := range a {
		fmt.Printf("indice: %d, valor: %d\n", index, value)
	}

	var matrix = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)

}

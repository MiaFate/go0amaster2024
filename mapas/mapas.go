package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	fmt.Println(colors)
	fmt.Println(colors["rojo"])
	colors["negro"] = "#000000"
	fmt.Println(colors)
	valor, ok := colors["rojo"]
	if ok {
		fmt.Println("valor:", valor)
	} else {
		fmt.Println("no existe la clave")
	}
	delete(colors, "azul")
	fmt.Println(colors)

	fmt.Println("==== iterando mapa ====")
	for k, v := range colors {
		fmt.Printf("clave: %s, valor: %s\n", k, v)
	}
}

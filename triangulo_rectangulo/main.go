package main

import (
	"fmt"
	"math"
)

const Presition = 2

var a, b float64

func inicializar(_a float64, _b float64) {
	a = _a
	b = _b
}

// solicitar lados de un triangulo rectangulo por pantalla
func getSides() {
	fmt.Print("ingrese el lado a: ")
	_, err := fmt.Scan(&a)
	for err != nil {
		fmt.Print("error de ingreso, debe ingresar un numero: ")
		_, err = fmt.Scan(&a)
	}
	fmt.Print("ingrese el lado b: ")
	_, err2 := fmt.Scan(&b)
	for err2 != nil {
		fmt.Print("error de ingreso, debe ingresar un numero: ")
		_, err2 = fmt.Scan(&b)
	}

}

// Calcular Hipotenusa
func calcHip() float64 {
	hip2 := a*a + b*b
	hip := math.Sqrt(hip2)
	return hip
}

// Calcular Area
func calcArea() float64 {
	area := (a * b) / 2
	return area
}

// Calcular Perimetro
func calcPerimeter() float64 {
	hip := calcHip()
	perimeter := a + b + hip
	return perimeter
}

func main() {
	getSides()

	perimeter := calcPerimeter()
	area := calcArea()

	fmt.Printf("el perimetro es: %.*f\n", Presition, perimeter)
	fmt.Printf("el area es: %.*f\n", Presition, area)
}

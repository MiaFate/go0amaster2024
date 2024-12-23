// paquete principal del programa, aqui buscara la funcion main
package main

// Paquetes importados, si no los tenemos se pueden descargar con 'go get <package>'
import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de contantes
// local con minusculas
const pi = 3.14

const (
	x = 100
	y = 0b1010 // binario
	z = 0o12   // hexadecimal
	w = 0xFF   // octal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

// Global con mayusculas
const Pi = 3.1415

// Funcion principal, por aqui comienza nuestro programa
func main() {
	fmt.Println("hola mundo")
	fmt.Println(quote.Go())

	// Declaraciones de variables (con var no hace falta inicializarlas)
	var (
		firstName, lastName = "Mia", "Ramos"
		age                 = 40
	)

	// Delaracion y asignacion en linea (es obligatorio inicializarlas, solo se puede usar dentro de una  funcion)
	a, b, c := "Mia", "Ramos", 40

	// Impresion por consola
	fmt.Println(a, b, c)
	fmt.Println(firstName, lastName, age)
	fmt.Println(pi)

	fmt.Println(x, y, z, w)

	fmt.Println(Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)
}

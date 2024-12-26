package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	var p Person
	p.Name = "Mia"
	p.Age = 40
	p.Email = "mia@gmail.com"
	fmt.Println(p)

	mia := Person{
		"Mia",
		40,
		"mia@gmail.com",
	}
	fmt.Println(mia)

	fmt.Printf("el email de %s es: %s\n", mia.Name, mia.Email)
	mia.Age = 41
	fmt.Println("este año cumple:", mia.Age)

}

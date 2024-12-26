package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}

func (p *Person) sayHello() {
	fmt.Printf("Hello, my name is %s!\n", p.name)
}

func main() {
	var x int = 10
	//var p *int = &x
	//fmt.Println(x)
	//fmt.Println(*p)
	//x++
	//fmt.Println(x)
	//fmt.Println(*p)
	fmt.Println(x)
	edit(&x)
	fmt.Println(x)

	p := Person{"Mia", 40, "mia@gmail.com"}
	p.sayHello()
}

func edit(x *int) {
	*x = 20
}

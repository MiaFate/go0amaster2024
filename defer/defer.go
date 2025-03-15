package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("contact-manager/defer/hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola mundo!"))
	if err != nil {
		fmt.Println(err)
		return
	}

}

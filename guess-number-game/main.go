package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func play() {
	const maxTries = 10
	randomNum := rand.Intn(100)
	var enteredNum int
	var tries int

	for tries < maxTries {
		tries++
		fmt.Printf("Ingrese un numero (Intentos restantes: %d): ", maxTries-tries+1)
		_, _ = fmt.Scanln(&enteredNum)

		if enteredNum == randomNum {
			fmt.Println("Felicitacions, adivinaste el numero!")
			playAgain()
		} else if enteredNum > randomNum {
			fmt.Println("El numero a adivinar es menor")
		} else if enteredNum < randomNum {
			fmt.Println("El numero a adivinar es mayor")
		}
	}

	fmt.Println("Se acabaron los intentos, el numero era: ", randomNum)
	playAgain()
}

func playAgain() {
	var option string
	fmt.Print("Queres jugar de nuevo? (s/n): ")
	_, err := fmt.Scanln(&option)
	if err != nil {
		return
	}
	switch option {
	case "s":
		play()
	case "n":
		fmt.Println("Gracias por jugar!")
	default:
		fmt.Println("eleccion invalida!")
		playAgain()
	}
}

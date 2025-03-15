package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// estructura de contactos
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guarda contactos en un archivo json
func saveContacts(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// slice de contactos
	var contacts []Contact
	// carga contactos desde archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error cargando contactos desde archivo: ", err)
	}

	// crear instancia de buffio
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("==== Contact Manager ====\n",
			"1. Agregar contacto\n",
			"2. Mostrar contactos\n",
			"3. Salir\n",
			"Opción: ")

		// leer opcion del usuario
		var opcion int
		_, err = fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error leyendo la opción: ", err)
			return
		}

		//manejar la opcion del usuario
		switch opcion {
		case 1:
			// agregar contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')

			// agregar contacto al slice
			contacts = append(contacts, c)

			// guardar contactos en archivo
			if err = saveContacts(contacts); err != nil {
				fmt.Println("Error guardando contactos en archivo: ", err)
			}
		case 2:
			// mostrar contactos
			fmt.Println("==== Contactos ====")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("==========================")
		case 3:
			// salir
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}

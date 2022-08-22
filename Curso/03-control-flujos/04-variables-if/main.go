package main

import (
	"fmt"
)

func main() {

	if nombre, edad := "Flavio", 26; nombre == "Tomas" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	users := make(map[string]string)

	users["Flavio"] = "villanuevaflavioT@gmail.con"
	users["Tomas"] = "tomas@gmail.con"


	if correo, ok := users["Roel"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}

}

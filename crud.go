package main

import (
	"fmt"
	"log"
)

func main() {
	e := Estudiante{
		Name:   "Aron",
		Age:    22,
		Active: true,
	}

	err := EstudianteCrear(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creado exitosamente!!")
}

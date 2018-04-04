package main

import (
	"fmt"
)

func main() {
	/* e := Estudiante{
		Name:   "Aron",
		Age:    22,
		Active: true,
	}

	err := EstudianteCrear(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creado exitosamente!!") */

	es, err := ConsultarEstudiante()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(es)
}

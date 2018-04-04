package main

import "fmt"

func main() {

	// CREAR ESTUDIANTE
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

	// CONSULTAR ESTUDIANTE
	/* es, err := ConsultarEstudiante()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(es) */

	//ACTUALIZAR estudiante
	/* e := Estudiante{
		ID:     3,
		Name:   "pedro",
		Age:    44,
		Active: true,
	}
	err := ActualizarEstudiante(e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Actualizado correctamente") */

	//BORRAR estudiante

	err := BorrarEstudiante(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Borrado correctamente")
	}
}

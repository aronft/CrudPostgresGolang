package main

import (
	"errors"
	"time"
)

//Estudiante estructura del estudiante
type Estudiante struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//EstudianteCrear registra un estudiante en la BD
func EstudianteCrear(e Estudiante) error {
	// Primero crear el string con la instruccion que se va a ejecutar
	q := `INSERT INTO 
			estudiantes (name, age, active)
			VALUES ($1, $2, $3)`
	db := getConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(e.Name, e.Age, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}

	return nil
}

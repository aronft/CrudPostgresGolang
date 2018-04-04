package main

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"

	_ "github.com/lib/pq"
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

	intNull := sql.NullInt64{}
	strNull := sql.NullString{}
	// boolNull := sql.NullBool{}

	db := getConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if e.Age == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(e.Age)
	}

	if e.Name == "" {
		strNull.Valid = false
	} else {
		strNull.Valid = true
		strNull.String = e.Name
	}

	r, err := stmt.Exec(strNull, intNull, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}

	return nil
}

//ConsultarEstudiante busca la inormación de los estudiantes (todos)
func ConsultarEstudiante() (estudiantes []Estudiante, err error) {
	q := `SELECT id, name, age, active, created_at, updated_at
			FROM estudiantes;`
	// En caso que se tena un tipo de dato repetido se debera crear una variable para cada estudiante
	timeNUll := pq.NullTime{}
	intNUll := sql.NullInt64{}
	strNUll := sql.NullString{}
	// boolNull := sql.NullBool{}

	db := getConection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := Estudiante{}
		err = rows.Scan(
			&e.ID,
			&strNUll,
			&intNUll,
			&e.Active,
			&e.CreatedAt,
			&timeNUll,
		)
		if err != nil {
			return
		}
		e.UpdatedAt = timeNUll.Time
		e.Name = strNUll.String
		e.Age = int16(intNUll.Int64)
		estudiantes = append(estudiantes, e)
	}

	return estudiantes, nil
}

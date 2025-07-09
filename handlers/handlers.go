package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// listar los socios suscrito, data viene desde la base de datos
func ListContacts(db *sql.DB) {
	//
	query := "SELECT id, nombres, apellido_paterno, apellido_materno, correo, celular FROM suscripcion"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------")
	for rows.Next() {
		// INSTANCIA DE MODELO SUSCRIPCION
		suscripcion := models.Suscripcion{}
		var valueNombres, valueApellidoPaterno, valueApellidoMaterno, valueCorreo, valueCelular sql.NullString
		err := rows.Scan(&suscripcion.Id, &valueNombres, &valueApellidoPaterno, &valueApellidoMaterno, &valueCorreo, &valueCelular)
		if err != nil {
			log.Fatal(err)
		}
		if valueNombres.Valid {
			suscripcion.Nombres = valueNombres.String
		} else {
			suscripcion.Nombres = "Sin Nombres"
		}
		if valueApellidoPaterno.Valid {
			suscripcion.ApellidoPaterno = valueApellidoPaterno.String
		} else {
			suscripcion.ApellidoPaterno = "Sin Apellido Paterno"
		}
		if valueApellidoMaterno.Valid {
			suscripcion.ApellidoMaterno = valueApellidoMaterno.String
		} else {
			suscripcion.ApellidoMaterno = "Sin Apellido Materno"
		}
		if valueCorreo.Valid {
			suscripcion.Correo = valueCorreo.String
		} else {
			suscripcion.Correo = "Sin Correo"
		}
		if valueCelular.Valid {
			suscripcion.Celular = valueCelular.String
		} else {
			suscripcion.Celular = "Sin Celular"
		}

		fmt.Printf("ID: %d || Nombres: %s  Apellido Paterno: %s  Apellido Materno: %s  Correo: %s  Celular: %s \n",
			suscripcion.Id, suscripcion.Nombres, suscripcion.ApellidoPaterno, suscripcion.ApellidoMaterno, suscripcion.Correo, suscripcion.Celular)
		fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------")
	}
}

// GetSuscripcionById es una metodo para buscar un registro atraves de un id
func GetSuscripcionById(db *sql.DB, id int) {
	query := "SELECT id, nombres, apellido_paterno, apellido_materno, correo, celular FROM suscripcion WHERE id = ? "

	row := db.QueryRow(query, id)
	suscripcion := models.Suscripcion{}
	var valueNombres, valueApellidoPaterno, valueApellidoMaterno, valueCorreo, valueCelular sql.NullString
	err := row.Scan(&suscripcion.Id, &valueNombres, &valueApellidoPaterno, &valueApellidoMaterno, &valueCorreo, &valueCelular)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("\nRESULTADO ")
			fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------")
			log.Fatalf("No se encontro ningun registro de suscripcion con el ID: %d \n", id)
		}
	}
	if valueNombres.Valid {
		suscripcion.Nombres = valueNombres.String
	} else {
		suscripcion.Nombres = "Sin Nombres"
	}
	if valueApellidoPaterno.Valid {
		suscripcion.ApellidoPaterno = valueApellidoPaterno.String
	} else {
		suscripcion.ApellidoPaterno = "Sin Apellido Paterno"
	}
	if valueApellidoMaterno.Valid {
		suscripcion.ApellidoMaterno = valueApellidoMaterno.String
	} else {
		suscripcion.ApellidoMaterno = "Sin Apellido Materno"
	}
	if valueCorreo.Valid {
		suscripcion.Correo = valueCorreo.String
	} else {
		suscripcion.Correo = "Sin Correo"
	}
	if valueCelular.Valid {
		suscripcion.Celular = valueCelular.String
	} else {
		suscripcion.Celular = "Sin Celular"
	}

	fmt.Printf("\n SUSCRIPCION: %d \n", id)
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------")

	fmt.Printf("Nombres: %s || Apellido Paterno: %s || Apellido Materno: %s || Correo: %s || Celular: %s \n",
		suscripcion.Nombres, suscripcion.ApellidoPaterno, suscripcion.ApellidoMaterno, suscripcion.Correo, suscripcion.Celular)
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------")

}

func CreateContact(db *sql.DB, suscripcion models.Suscripcion) {

	query := "INSERT INTO suscripcion ( nombres, apellido_paterno, apellido_materno, correo, celular) VALUES( ? , ? , ? , ? , ?);"
	_, err := db.Exec(query, suscripcion.Nombres, suscripcion.ApellidoPaterno, suscripcion.ApellidoMaterno, suscripcion.Correo, suscripcion.Celular)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("El registro fue guardado correctamente...!!!")
	ListContacts(db)

}

// vamos a crear la funcion para actualizar una suscripcion
func UpdateSuscripcion(db *sql.DB, suscripcion models.Suscripcion) {

	query := "UPDATE suscripcion SET nombres=?, apellido_paterno=?, apellido_materno=?, correo=?, celular=? WHERE id= ? ;"
	_, err := db.Exec(query, suscripcion.Nombres, suscripcion.ApellidoPaterno, suscripcion.ApellidoMaterno, suscripcion.Correo, suscripcion.Celular, suscripcion.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Se actualizo el registro fue guardado correctamente...!!!")
	ListContacts(db)
}

// vamos a crear la funcion que elimina un regitro de la base de datos por id
func DeleteSuscripcion(db *sql.DB, id int) {

	query := "DELETE FROM suscripcion WHERE id= ?;"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Se elimino el registro %d correctamente...!!!", id)
	ListContacts(db)

}

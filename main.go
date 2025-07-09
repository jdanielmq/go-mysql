package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("MySql and Go")
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Conexion a la base de datos MySql Exitosa...!!!")

	handlers.ListContacts(db)
	handlers.GetSuscripcionById(db, 2)
	//handlers.GetSuscripcionById(db, 3)

	suscripcion := models.Suscripcion{
		Nombres:         "Ignacio Zuñiga",
		ApellidoPaterno: "Zuñiga",
		ApellidoMaterno: "Diaz",
		Correo:          "ignacio.zuniga.diaz@gmail.com",
		Celular:         "+56965430987",
	}
	fmt.Printf("\n\n\n")
	log.Println("Iniciando el ingreso de registro!!!")
	handlers.CreateContact(db, suscripcion)

	suscripcionUpdate := models.Suscripcion{
		Id:              1,
		Nombres:         "JUAN DANIEL",
		ApellidoPaterno: "MUÑOZ",
		ApellidoMaterno: "QUEUPUL",
		Correo:          "JDANIELMQ@GMAIL.COM",
		Celular:         "+3456789854",
	}
	fmt.Printf("\n\n\n")
	log.Printf("Iniciando actualización del registro %d", suscripcionUpdate.Id)
	handlers.UpdateSuscripcion(db, suscripcionUpdate)

	idSuscripcion := 9
	fmt.Printf("\n\n\n")
	log.Printf("Iniciando eliminacion del registro %d", idSuscripcion)
	//handlers.DeleteSuscripcion(db, idSuscripcion)

	/* dns := "jdanielmq:JdmQ1481@tcp(localhost:3306)/db_grossgym"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}*/
	for {
		fmt.Println("\n Menú")
		fmt.Println("1.- Listar Suscripciones")
		fmt.Println("2.- Obtener Suscripción por ID")
		fmt.Println("3.- Crear nuevo Suscripción")
		fmt.Println("4.- Actualizar Suscripción")
		fmt.Println("5.- Eliminar Suscripción")
		fmt.Println("6.- Salir")
		fmt.Print("Seleccione una opción: ")

		// leer las opcion seleccionada por el usuario
		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("ingrese el id de la suscripcion: ")
			var id int
			fmt.Scanln(&id)
			handlers.GetSuscripcionById(db, id)
		case 3:
			suscripcion := inputSuscripcionDetails(3)
			handlers.CreateContact(db, suscripcion)
		case 4:
			suscripcionUpdate := inputSuscripcionDetails(4)
			handlers.UpdateSuscripcion(db, suscripcionUpdate)
		case 5:
			fmt.Print("ingrese el id de la suscripcion a eliminar: ")
			var id int
			fmt.Scanln(&id)
			handlers.DeleteSuscripcion(db, id)
		case 6:
			fmt.Println("Saliendo del Programa...")
			return
		default:
			fmt.Println("opcion no valida, verifique en el menu.")
		}
	}

}

func inputSuscripcionDetails(opcion int) models.Suscripcion {
	reader := bufio.NewReader(os.Stdin)

	var suscripcion models.Suscripcion
	if opcion == 4 {
		fmt.Print("ingrese el id de la suscripcion a editar: ")
		var id int
		fmt.Scanln(&id)
		suscripcion.Id = id
	}

	fmt.Print("Ingrese el nombre del suscrito: ")
	name, _ := reader.ReadString('\n')
	suscripcion.Nombres = strings.TrimSpace(name)

	fmt.Print("Ingrese el apellido paterno del suscrito: ")
	paterno, _ := reader.ReadString('\n')
	suscripcion.ApellidoPaterno = strings.TrimSpace(paterno)

	fmt.Print("Ingrese el apellido materno del suscrito: ")
	materno, _ := reader.ReadString('\n')
	suscripcion.ApellidoMaterno = strings.TrimSpace(materno)

	fmt.Print("Ingrese el correo del suscrito: ")
	correo, _ := reader.ReadString('\n')
	suscripcion.Correo = strings.TrimSpace(correo)

	fmt.Print("Ingrese el celular del suscrito: ")
	celular, _ := reader.ReadString('\n')
	suscripcion.Celular = strings.TrimSpace(celular)

	return suscripcion

}

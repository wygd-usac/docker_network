package db

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/Pallinder/go-randomdata"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

var nombre_servidor = "default"
var mensaje = "Mensaje default"

var db_handler *sql.DB;
//PruebaHandler funcion
func PruebaHandler (w http.ResponseWriter, request *http.Request) {
	enableCors(&w)

	w.Write([]byte("Yo no hago nada, pero soy accesible desde el ingress"))
}

//MensajeHandler funcion
func MensajeHandler (w http.ResponseWriter, request *http.Request) {
	enableCors(&w)

	w.Write([]byte(mensaje))
}

func DatosHandler (w http.ResponseWriter, request *http.Request) {
	enableCors(&w)
	results, err := db_handler.Query("SELECT EmployeeId, FirstName FROM EMPLOYEE")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	registros:= ""
    for results.Next() {
        var emp Employee
        // for each row, scan the result into employee object
        err = results.Scan(&emp.Id, &emp.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        registros = registros + strconv.Itoa(emp.Id) + " - " + emp.Name + "\n"
    }
    registros = registros + mensaje
	w.Write([]byte(registros))
}

func main(){
	router := mux.NewRouter()


	router.HandleFunc("/url_prueba", PruebaHandler).Methods("GET")
	router.HandleFunc("/mensaje",MensajeHandler).Methods("GET")
	router.HandleFunc("/datos",DatosHandler).Methods("GET")
	nombre_servidor = randomdata.SillyName()
	mensaje = "Hola! te saluda el servidor " + nombre_servidor
	log.Println("Estoy escuchando en el puerto 5000")
	
	
	//Se utilizan variables de entorno para poder modificar valores desde el archivo docker-compose, y 
	//asi evitar tener que modificar el codigo, y generar nuevamente la imagen.
	usuario:= os.Getenv("USER_NAME")
	pass:= os.Getenv("PASS")
	host:= os.Getenv("HOST")
	port:= os.Getenv("PORT")
	db_name:= os.Getenv("DB_NAME")
	conn_string := usuario+":"+pass+"@tcp("+host+":"+port+")/"+db_name
	db, err := sql.Open("mysql", conn_string)

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }else{
		db_handler = db
		log.Println("conectado a la base de datos")
		}
	

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
	
	log.Fatal(http.ListenAndServe(":5000", router))
	
	

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

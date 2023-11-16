package main

import (
	"net/http"
	"log"
	//"fmt"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main()  {
	http.HandleFunc("/", Home)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprintf(w, "Hola iory")
	plantillas.ExecuteTemplate(w, "home", nil)
}
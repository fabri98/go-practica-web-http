package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// Definimos un manejador de archivos para servir archivos estáticos desde el directorio "./ui/static/"
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Usamos http.StripPrefix para eliminar el prefijo "/static" de la URL
	// antes de pasarla al manejador de archivos.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("POST /tasks/create", CreateTask)
	mux.HandleFunc("GET /tasks", GetTasks)
	// En este caso, el id se pasa como un parámetro en la URL, por ejemplo: /tasks/delete/1
	mux.HandleFunc("DELETE /tasks/delete/{id}", DeleteTaskById)
	// En este caso, el id se pasa como un parámetro de consulta en la URL, por ejemplo: /tasks/find?id=1
	mux.HandleFunc("GET /tasks/find", GetTaskById)

	fmt.Println("Listen...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

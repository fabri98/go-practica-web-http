package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	// Definimos los archivos de plantilla que se van a utilizar para renderizar la página
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}
	// Cargar la plantilla HTML desde un archivo
	template, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Renderizar la plantilla y enviarla como respuesta.
	// El segundo parámetro es el objeto de datos que se pasará a la plantilla.
	err = template.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating task")
	//De esta forma podemos enviar un json de respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"Saludar": "Hola"}`))

}
func GetTasks(w http.ResponseWriter, r *http.Request) {
	// De esta forma podemos enviar un json de respuesta al cliente
	w.Header().Set("Cache-Control", "max-age=3600")

	// Agregar un valor a un encabezado existente
	w.Header().Add("Cache-Control", "public")
	// Eliminar todos los valores de un encabezado
	w.Header().Del("Cache-Control")
	// Eliminar la fecha de respuesta del encabezado
	w.Header()["Date"] = nil
	fmt.Println("Get all tasks")
}
func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("Delete")

	fmt.Fprintf(w, "El id ha eliminar es: %s", id)
}
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

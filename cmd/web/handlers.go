package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// Definimos los archivos de plantilla que se van a utilizar para renderizar la página
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}
	// Cargar la plantilla HTML desde un archivo
	template, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
		return
	}
	// Renderizar la plantilla y enviarla como respuesta.
	// El segundo parámetro es el objeto de datos que se pasará a la plantilla.
	err = template.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
		return
	}
}
func (app *application) CreateTask(w http.ResponseWriter, r *http.Request) {
	//De esta forma podemos enviar un json de respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Task created successfully"}`))

}

func (app *application) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

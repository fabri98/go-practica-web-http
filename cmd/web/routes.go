package main

import "net/http"

// El metodo routes() define las rutas de la aplicación y
// devuelve un multiplexer de rutas (ServeMux) que se utilizará para manejar las solicitudes HTTP entrantes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Definimos un manejador de archivos para servir archivos estáticos desde el directorio "./ui/static/"
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Usamos http.StripPrefix para eliminar el prefijo "/static" de la URL
	// antes de pasarla al manejador de archivos.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /tasks/create", app.CreateTask)
	// En este caso, el id se pasa como un parámetro de consulta en la URL, por ejemplo: /tasks/find?id=1
	mux.HandleFunc("GET /tasks/find", app.GetTaskById)

	return mux
}

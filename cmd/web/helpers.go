package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError registra un error en el registro de errores (stack trace) y
// envía una respuesta de error 500 al cliente.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError envía una respuesta de error al cliente con el código de estado especificado.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound envía una respuesta de error 404 al cliente. Utiliza un wrapper
// para llamar a clientError con el código de estado 404.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

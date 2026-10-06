package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	// Inicializamos la bandera de línea de comandos para especificar la dirección de red
	// en la que el servidor escuchará las solicitudes HTTP.
	// El valor por defecto que definimos es ":8080", lo que significa que escuchará en el puerto 8080
	// en todas las interfaces de red.
	addr := flag.String("addr", ":8080", "HTTP network address")

	// Analizamos los argumentos de la línea de comandos para obtener el valor de la flag "addr".
	flag.Parse()

	// Creamos dos loggers: uno para mensajes de información y otro para mensajes de error.
	// El logger de información escribirá en la salida estándar (os.Stdout)
	// y el logger de error escribirá en la salida de error estándar (os.Stderr).
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creamos una instancia de la aplicación y le pasamos los loggers.
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// Creamos un server HTTP personalizado con la dirección de red,
	// el logger de errores y el multiplexer de rutas.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

package main

import (
	"log"
	"net/http"
	r "rpsweb/handlers"
)

func main() {
	//crear enrutador
	router := http.NewServeMux()
	//manejador de archivos estaticos desde una direccion
	fs:=http.FileServer(http.Dir("static"))
	//ruta para acceder a los archivos estaticos
	router.Handle("/static/",http.StripPrefix("/static/",fs))
	//configurar rutas
	router.HandleFunc("/", r.Index)
	router.HandleFunc("/new-game", r.NewGame)
	router.HandleFunc("/game", r.Game)
	router.HandleFunc("/jugar",r.Play)
	router.HandleFunc("/about", r.About)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
	
}

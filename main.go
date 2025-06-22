package main

import (
	"log"
	"net/http"
	r "rpsweb/handlers"
)

func main() {
	//crear enrutador
	router := http.NewServeMux()
	//configurar rutas
	router.HandleFunc("/", r.Index)
	router.HandleFunc("/new", r.NewGame)
	router.HandleFunc("/game", r.Game)
	router.HandleFunc("/play", r.Play)
	router.HandleFunc("/about", r.About)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
	
}

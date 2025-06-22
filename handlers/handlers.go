package handlers

import (
	"fmt"
	"net/http"
)

// manejador
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hola mundo,desde Roel</h1>")
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juego")
}
func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}
func Play(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Jugar")
}
func About(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Acerca de...")

}
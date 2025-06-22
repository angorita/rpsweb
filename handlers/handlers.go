package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// manejador
func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar", http.StatusInternalServerError)
		return
	}
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juego")
}
func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}
func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Jugar")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Acerca de...")

}

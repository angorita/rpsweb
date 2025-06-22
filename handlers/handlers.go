package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

const (
	templateDir  = "template/"
	templateBase = templateDir + "base.html"
)

// para enviar datos a una plantilla en go necesitamos una estructura,mapas u otros
// manejador
func Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html",nil)
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
func RenderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar", http.StatusInternalServerError)
log.Println(err)
		return
	}
}

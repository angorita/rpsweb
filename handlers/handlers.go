package handlers

import (
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "template/"
	templateBase = templateDir + "base.html"
)

// para enviar datos a una plantilla en go necesitamos una estructura,mapas u otros
// manejador
func Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "new-game.html", nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "game.html", nil)
}
func Play(w http.ResponseWriter, r *http.Request){
	RenderTemplate(w,"jugar.html",nil)
}
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html", nil)
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

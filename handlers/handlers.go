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

type Player struct {
	Name string
}

var player Player

// para enviar datos a una plantilla en go necesitamos una estructura,mapas u otros
// manejador
func Index(w http.ResponseWriter, r *http.Request) {
	RestartValue()
	RenderTemplate(w, "index.html", nil)
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	RestartValue()
	RenderTemplate(w, "new-game.html", nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error Parsing form", http.StatusBadRequest)
			return
		}
		//capturo name de new-game.html
		player.Name = r.Form.Get("name")
	}
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}
	RenderTemplate(w, "new-game.html", player)
}
func Play(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "jugar.html", nil)
}
func About(w http.ResponseWriter, r *http.Request) {
	RestartValue()
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
//reiniciar valores
func RestartValue(){
	player.Name=""

}

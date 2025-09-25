package controller

import (
	"net/http"
	"power4/controller/structure"
	"power4/controller/utils"
	"power4/pages"
	"sync"
)

func renderPage(w http.ResponseWriter, filename string, data any) {
	err := pages.Temp.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Erreur rendu template : "+err.Error(), http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "Accueil",
		"Message": "Bienvenue sur la page d'accueil !",
	}
	renderPage(w, "index.html", data)
}

func About(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "A propos",
		"Message": "Ceci est la page a propos",
	}
	renderPage(w, "about.html", data)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		msg := r.FormValue("msg")

		data := map[string]string{
			"Title":   "Contact",
			"Message": "Merci " + name + " pour ton message " + msg,
		}

		renderPage(w, "contact.html", data)
		return
	}
	data := map[string]string{
		"Title":   "Contact",
		"Message": "Rentrer votre message",
	}
	renderPage(w, "contact.html", data)

}

var (
	mu            sync.Mutex
	currentPlayer = 1 // 1 = rouge, 2 = jaune
)

func Step(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		data := map[string]string{
			"Title":   "Jeu",
			"Message": "Choisis une pièce pour commencer",
		}
		renderPage(w, "index.html", data)
		return
	}

	choice := r.FormValue("piece") // assure-toi que tes boutons ont name="piece" et value="..."

	// section critique : lire/modifier currentPlayer
	mu.Lock()
	var color string
	if currentPlayer == 1 {
		color = "red"
	} else {
		color = "yellow"
	}

	table := utils.PlacePiece(choice, color, &structure.Table{})
	data := map[string]any{
		"Title":      "Jeu",
		"Message":    "C'est au joueur " + color + " de jouer. Tu as choisi la pièce " + choice,
		"Placements": table,
	}
	// alterner le joueur pour le prochain tour
	if currentPlayer == 1 {
		currentPlayer = 2
	} else {
		currentPlayer = 1
	}
	mu.Unlock()

	renderPage(w, "index.html", data)
}

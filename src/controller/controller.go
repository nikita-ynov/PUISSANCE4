package controller

import (
	"fmt"
	"html/template"
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
	data := map[string]any{
		"Title":      "Accueil",
		"Message":    "Bienvenue sur la page d'accueil !",
		"Placements": gameTable.Placement,
		"redName":    redName,
		"yellowName": yellowName,
		"Form":       true,
	}
	renderPage(w, "index.html", data)
}

func About(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "A propos",
		"Message": "Ceci est la page à propos",
	}
	renderPage(w, "about.html", data)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		msg := r.FormValue("msg")

		data := map[string]string{
			"Title":   "Contact",
			"Message": "Merci " + name + " pour ton message : " + msg,
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
	currentPlayer = 1                  // 1 = rouge, 2 = jaune
	gameTable     = &structure.Table{} // état global
	redName       = "Rouge"
	yellowName    = "Jaune"
)

func ChangeName(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		rName := r.FormValue("redName")
		yName := r.FormValue("yellowName")

		if len(rName) >= 1 {
			redName = rName
		}

		if len(yName) >= 1 {
			yellowName = yName
		}

		data := map[string]any{
			"Title":      "Accueil",
			"Message":    template.HTML("Vous avez bien changé les noms! C'est le jouer <span class='red'> " + redName + "</span> qui commance."),
			"Placements": gameTable.Placement,
			"redName":    redName,
			"yellowName": yellowName,
			"Form":       false,
		}
		fmt.Print(data)
		renderPage(w, "index.html", data)
		return
	}
}

func Step(w http.ResponseWriter, r *http.Request) {
	winner := utils.CheckPlacement(gameTable)

	if winner != "" {
		// On annonce le gagnant et on propose de rejouer via un refresh/accueil.
		text := redName
		if winner == "yellow" {
			text = yellowName
		}
		data := map[string]any{
			"Title":      "Jeu",
			"Message":    template.HTML("Le joueur <span class='" + winner + "'>" + text + "</span> a gagné !"),
			"Placements": gameTable.Placement,
			"redName":    redName,
			"yellowName": yellowName,
			"Winner":     winner,
		}
		renderPage(w, "index.html", data)
		return
	} else {
		if r.Method != http.MethodPost {
			renderPage(w, "index.html", map[string]any{
				"Title":      "Jeu",
				"Message":    "Choisis une pièce pour commencer",
				"redName":    redName,
				"yellowName": yellowName,
				"Placements": gameTable.Placement,
			})
			return
		}

		choice := r.FormValue("piece")

		mu.Lock()
		defer mu.Unlock()

		// Couleur du joueur courant
		color := "red"
		if currentPlayer == 2 {
			color = "yellow"
		}

		// 1) Essayer de placer la pièce
		var placed bool
		gameTable, placed = utils.PlacePiece(choice, color, gameTable)

		// Si la pose échoue (colonne pleine), ne pas changer de joueur
		if !placed {
			data := map[string]any{
				"Title":      "Jeu",
				"Message":    "Colonne pleine ! Choisis une autre colonne.",
				"redName":    redName,
				"yellowName": yellowName,
				"Placements": gameTable.Placement,
			}
			renderPage(w, "index.html", data)
			return
		}

		// 2) Vérifier s'il y a un gagnant
		winner := utils.CheckPlacement(gameTable)

		if winner != "" {
			text := redName
			if winner == "yellow" {
				text = yellowName
			}
			data := map[string]any{
				"Title":      "Jeu",
				"Message":    template.HTML("Le joueur <span class='" + winner + "'>" + text + "</span> a gagné !"),
				"Placements": gameTable.Placement,
				"redName":    redName,
				"yellowName": yellowName,
				"Winner":     winner,
			}
			renderPage(w, "index.html", data)
			return
		}

		// 3) Alterner le joueur
		if currentPlayer == 1 {
			currentPlayer = 2
		} else {
			currentPlayer = 1
		}

		// Couleur du prochain joueur (pour le message)
		nextColor := "red"
		text := redName
		if currentPlayer == 2 {
			nextColor = "yellow"
			text = yellowName
		}

		data := map[string]any{
			"Title":      "Jeu",
			"Message":    template.HTML("Tu as joué la pièce " + choice + ". À <span class='" + nextColor + "'>" + text + "</span> de jouer."),
			"Placements": gameTable.Placement,
			"redName":    redName,
			"yellowName": yellowName,
			"color":      nextColor,
		}
		renderPage(w, "index.html", data)
	}
}

func Reset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Réinitialiser l'état global
	gameTable = &structure.Table{}
	currentPlayer = 1

	// Retour à l'accueil propre
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

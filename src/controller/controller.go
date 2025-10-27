package controller

import (
	"html/template"
	"net/http"
	"power4/controller/structure"
	"power4/controller/utils"
	"power4/pages"
	"sync"
)

// BD temporaire des scores
var scores []structure.Score
var nextID = 1

// Verrou de fin de partie
var gameFinished = false

func renderPage(w http.ResponseWriter, filename string, data any) {
	err := pages.Temp.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Erreur rendu template : "+err.Error(), http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":        "Accueil",
		"Message":      "Bienvenue sur la page d'accueil !",
		"Placements":   gameTable.Placement,
		"redName":      redName,
		"yellowName":   yellowName,
		"Form":         true,
		"GameFinished": gameFinished,
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
			"Title":        "Accueil",
			"Message":      template.HTML("Vous avez bien changé les noms! C'est le jouer <span class='red'> " + redName + "</span> qui commance."),
			"Placements":   gameTable.Placement,
			"redName":      redName,
			"yellowName":   yellowName,
			"Form":         false,
			"GameFinished": gameFinished,
		}
		renderPage(w, "index.html", data)
		return
	}
}

func Step(w http.ResponseWriter, r *http.Request) {
	winner := utils.CheckPlacement(gameTable)

	// Si la partie est terminée, on bloque toute action et on affiche un message clair
	if gameFinished {
		msg := template.HTML("Partie terminée. Cliquez sur « Rejouer » pour relancer.")
		text := ""
		if winner != "" {
			text = redName
			if winner == "yellow" {
				text = yellowName
			}
			msg = template.HTML("Partie terminée. <span class='" + winner + "'>" + template.HTMLEscapeString(text) + "</span> a déjà gagné. Cliquez « Rejouer ».")
		}

		renderPage(w, "index.html", map[string]any{
			"Title":        "Jeu",
			"Message":      msg,
			"Placements":   gameTable.Placement,
			"redName":      redName,
			"yellowName":   yellowName,
			"Winner":       winner,
			"GameFinished": true,
		})
		return
	}

	// Si quelqu'un a gagné sur l'état actuel, on enregistre une seule fois
	if winner != "" && !gameFinished {
		gameFinished = true

		text := redName
		if winner == "yellow" {
			text = yellowName
		}

		// Sauvegarde du score
		scores = append(scores, structure.Score{
			ID:           nextID,
			RedPlayer:    redName,
			YellowPlayer: yellowName,
			Winner:       text,
		})
		nextID++

		data := map[string]any{
			"Title":        "Jeu",
			"Message":      template.HTML("Le joueur <span class='" + winner + "'>" + text + "</span> a gagné !"),
			"Placements":   gameTable.Placement,
			"redName":      redName,
			"yellowName":   yellowName,
			"Winner":       winner,
			"GameFinished": true,
		}
		renderPage(w, "index.html", data)
		return
	}

	// Méthode GET : afficher la page de jeu
	if r.Method != http.MethodPost {
		msg := "Choisis une pièce pour commencer"
		if gameFinished {
			msg = "Partie terminée. Cliquez sur « Rejouer » pour relancer."
		}
		renderPage(w, "index.html", map[string]any{
			"Title":        "Jeu",
			"Message":      msg,
			"redName":      redName,
			"yellowName":   yellowName,
			"Placements":   gameTable.Placement,
			"GameFinished": gameFinished,
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
			"Title":        "Jeu",
			"Message":      "Colonne pleine ! Choisis une autre colonne.",
			"redName":      redName,
			"yellowName":   yellowName,
			"Placements":   gameTable.Placement,
			"GameFinished": gameFinished,
		}
		renderPage(w, "index.html", data)
		return
	}

	// 2) Vérifier s'il y a un gagnant
	winner = utils.CheckPlacement(gameTable)
	if winner != "" && !gameFinished {
		gameFinished = true

		text := redName
		if winner == "yellow" {
			text = yellowName
		}

		// Sauvegarde du score
		scores = append(scores, structure.Score{
			ID:           nextID,
			RedPlayer:    redName,
			YellowPlayer: yellowName,
			Winner:       text,
		})
		nextID++

		data := map[string]any{
			"Title":        "Jeu",
			"Message":      template.HTML("Le joueur <span class='" + winner + "'>" + text + "</span> a gagné !"),
			"Placements":   gameTable.Placement,
			"redName":      redName,
			"yellowName":   yellowName,
			"Winner":       winner,
			"GameFinished": true,
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
		"Title":        "Jeu",
		"Message":      template.HTML("Tu as joué la pièce " + choice + ". À <span class='" + nextColor + "'>" + text + "</span> de jouer."),
		"Placements":   gameTable.Placement,
		"redName":      redName,
		"yellowName":   yellowName,
		"color":        nextColor,
		"GameFinished": gameFinished,
	}
	renderPage(w, "index.html", data)
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
	gameFinished = false

	// Retour à l'accueil propre
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Scores(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":  "Scores",
		"Scores": scores,
	}
	renderPage(w, "scores.html", data)
}

package main

import (
	"fmt"
	"net/http"
	initTemp "power4/pages"
	"power4/router"
)

func main() {
	initTemp.Init()
	r := router.New()
	fmt.Println("Serveur demarré sur http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

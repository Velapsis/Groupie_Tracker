package main
import (
"fmt"
"net/http"
"html/template"
)
func main() {
// Handler de la page d'accueil
http.HandleFunc("/", homeHandler)
// Démarrer le serveur sur le port 8080
err := http.ListenAndServe(":8080", nil)
if err != nil {
fmt.Println("Erreur lors du démarrage du serveur:", err)
}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
//
}
package server

import (
	"encoding/json"
	"fmt"
	groupie "main/logic"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"text/template"
)

// CreateWebsite initialise les routes et démarre le serveur web
func CreateWebsite() {
	// Configuration des dossiers statiques
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Configuration des routes
	http.HandleFunc("/", MainMenu)
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.HandleFunc("/search", SearchAPIHandler)
	http.HandleFunc("/geocode", GeocodeHandler)

	OpenBrowser("http://localhost:8000")
	fmt.Println("Server listening on port http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

// OpenBrowser lance le navigateur par défaut selon le système d'exploitation
func OpenBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

// SearchAPIHandler gère les requêtes de recherche avec filtres
func SearchAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Récupération des paramètres de recherche et filtres
	query := r.URL.Query().Get("query")
	filters := map[string]string{
		"creationDateMin": r.URL.Query().Get("creationDateMin"),
		"creationDateMax": r.URL.Query().Get("creationDateMax"),
		"albumDateMin":    r.URL.Query().Get("albumDateMin"),
		"albumDateMax":    r.URL.Query().Get("albumDateMax"),
		"location":        r.URL.Query().Get("location"),
		"members":         r.URL.Query().Get("members"),
	}

	artists, err := groupie.GetArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	results := groupie.SearchArtistsWithFilters(artists, query, filters)
	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		http.Error(w, "Failed to encode results", http.StatusInternalServerError)
		return
	}
}

// MainMenu gère l'affichage de la page d'accueil
func MainMenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		t, _ := template.ParseFiles("templates/error.html")
		t.Execute(w, http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("templates/menu.html")
	if err != nil {
		http.Error(w, "500: internal server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, r)
}

// AboutPage gère l'affichage de la page "À propos"
func AboutPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about/" {
		t, _ := template.ParseFiles("templates/error.html")
		t.Execute(w, http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, "500: internal server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, r)
}

// IndexHandler gère l'affichage de la liste des artistes
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/index" {
		t, _ := template.ParseFiles("templates/error.html")
		t.Execute(w, http.StatusNotFound)
		return
	}

	artists, err := groupie.GetArtists()
	if err != nil {
		http.Error(w, "500: internal server error", http.StatusInternalServerError)
		return
	}

	// Gestion des requêtes AJAX
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists)
		return
	}

	// Affichage normal de la page
	t, err := template.ParseFiles("templates/mainpage.html")
	if err != nil {
		http.Error(w, "500: internal server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, artists)
}

// ArtistHandler gère l'affichage des détails d'un artiste spécifique
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/artist.html")
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}

		artist, err := groupie.GetArtists()
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}

		id, _ := strconv.Atoi(r.FormValue("id"))
		t.Execute(w, artist[id-1])
	} else {
		http.Error(w, "400: bad request.", http.StatusBadRequest)
	}
}

// GeocodeHandler gère la conversion des noms de lieux en coordonnées géographiques
func GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")

	// Appel à l'API OpenStreetMap pour obtenir les coordonnées
	resp, err := http.Get(fmt.Sprintf(
		"https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1",
		url.QueryEscape(location),
	))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

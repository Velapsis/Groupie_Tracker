package server

import (
	"encoding/json"
	"fmt"
	groupie "main/logic"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"text/template"
)

func CreateWebsite() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", MainMenu)
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.HandleFunc("/search", SearchAPIHandler)
	//http.HandleFunc("/search", SearchHandler)

	//OpenBrowser("http://localhost:8080")
	fmt.Println("Server listening on port http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
func OpenBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()

}
func SearchAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("query")
	filters := map[string]string{
		"creation_date":    "",
		"first_album_date": "",
		"location":         "",
	}

	artists, err := groupie.GetArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	results := groupie.SearchArtistsWithFilters(artists, query, filters)

	// S'assurer que l'encodage JSON est correct
	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		http.Error(w, "Failed to encode results", http.StatusInternalServerError)
		return
	}
}

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

	// Si c'est une requête AJAX
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists)
		return
	}

	// Pour l'affichage normal de la page
	t, err := template.ParseFiles("templates/mainpage.html")
	if err != nil {
		http.Error(w, "500: internal server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, artists)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		fmt.Println("url : ", r.URL.Path)

		fmt.Println("id : ", r.FormValue("id"))

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
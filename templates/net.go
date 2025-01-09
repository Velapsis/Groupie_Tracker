package server

import (
	"fmt"
	groupie "main/logic"
	"net/http"
	"os/exec"
	"runtime"
	"text/template"
)

func CreateWebsite() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", MainMenu)
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/artist/", ArtistHandler)

	OpenBrowser("http://localhost:8080")
	fmt.Println("Server listening on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		if r.URL.Path != "/index" {
			t, _ := template.ParseFiles("templates/error.html")
			t.Execute(w, http.StatusNotFound)
			return
		}

		t, err := template.ParseFiles("templates/mainpage.html")
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}
		artist, err := groupie.GetArtists()
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, artist)
	} else {
		http.Error(w, "400: Bad Request", http.StatusBadRequest)
		return
	}
}

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

		urlString := string(r.URL.Path)[8:]

		for i, v := range artist {
			if v.Name == urlString {
				t.Execute(w, artist[i])
				return
			}
		}

		t, _ = template.ParseFiles("templates/error.html")
		t.Execute(w, http.StatusNotFound)
		return
	} else {
		http.Error(w, "400: bad request.", http.StatusBadRequest)
	}

}

func LoadArtist(w http.ResponseWriter, r *http.Request) {
	groupie.GetArtists()
	IndexHandler(w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

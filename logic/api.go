// Package groupie gère l'interaction avec l'API Groupie Tracker
package groupie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetArtists récupère tous les artistes de l'API avec leurs informations complètes
func GetArtists() ([]Artist, error) {
	var artists []Artist
	
	// Récupère les données de base des artistes
	body, err := createBody("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("GetArtist error")
		return []Artist{}, err
	}

	// Désérialise les données JSON dans la structure temporaire
	err = json.Unmarshal(body, &temp)
	if err != nil {
		fmt.Println("GetArtist error")
		return []Artist{}, err
	}

	// Récupère les données de relations (concerts) pour tous les artistes
	relations, err := getRelations()
	if err != nil {
		fmt.Println("GetArtist error")
		return []Artist{}, err
	}

	// Combine les données des artistes avec leurs relations
	for v := range temp {
		// Reformate les dates dans les relations pour utiliser des points au lieu des tirets
		for key, dates := range relations[v].Relations {
			var modified []string
			for _, date := range dates {
				modified = append(modified, strings.ReplaceAll(date, "-", "."))
			}
			relations[v].Relations[key] = modified
		}

		// Crée un nouvel artiste avec toutes les informations combinées
		value := Artist{
			temp[v].Id,
			temp[v].Image,
			temp[v].Name,
			temp[v].Members,
			temp[v].CreationDate,
			strings.ReplaceAll(temp[v].FirstAlbum, "-", "."),
			relations[v].Relations,
		}

		artists = append(artists, value)
	}

	return artists, err
}

// getRelations récupère les informations de concerts pour tous les artistes
func getRelations() ([]Relation, error) {
	body, err := createBody("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("getRelations error")
		return []Relation{}, err
	}

	var temp struct {
		Relations []Relation `json:"index"`
	}

	err = json.Unmarshal(body, &temp)
	if err != nil {
		fmt.Println("getRelations error")
		return []Relation{}, err
	}

	// Reformate les données des relations pour une meilleure lisibilité
	for i := range temp.Relations {
		relations := make(map[string][]string, 1)

		for key, dates := range temp.Relations[i].Relations {
			var modDates []string
			var modKey string

			// Remplace les tirets et underscores par des espaces et virgules
			modKey = strings.ReplaceAll(key, "-", ", ")
			modKey = strings.ReplaceAll(modKey, "_", " ")

			// Reformate les dates
			for _, date := range dates {
				modDates = append(modDates, strings.ReplaceAll(date, "-", "."))
			}

			relations[modKey] = modDates
			temp.Relations[i].Relations = relations
		}
	}

	return temp.Relations, err
}

// createBody effectue une requête HTTP GET et retourne le corps de la réponse
func createBody(link string) ([]byte, error) {
	var body []byte
	response, err := http.Get(link)
	if err != nil {
		fmt.Println("createBody error")
		return body, err
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("createBody error")
		return body, err
	}

	return body, err
}
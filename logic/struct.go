package groupie

// Website représente la structure de base du site web
type Website struct {
	Template string
}

// Artist représente un artiste avec toutes ses informations
type Artist struct {
	Id           int                 `json:"id"`           // Identifiant unique de l'artiste
	Image        string              `json:"image"`        // URL de l'image de l'artiste
	Name         string              `json:"name"`         // Nom de l'artiste ou du groupe
	Members      []string            `json:"members"`      // Liste des membres du groupe
	CreationDate int                 `json:"creationDate"` // Année de création du groupe
	FirstAlbum   string              `json:"firstAlbum"`   // Date du premier album
	Relations    map[string][]string `json:"relations"`    // Map des lieux de concert et leurs dates
}

// Relation représente les informations de concert d'un artiste
type Relation struct {
	Id        int                 `json:"id"`             // Identifiant de l'artiste
	Relations map[string][]string `json:"datesLocations"` // Map des lieux et dates de concert
}

// Concert représente un concert spécifique
type Concert struct {
	ConcertDate     []string `json:"dates"`     // Dates du concert
	ConcertLocation []string `json:"locations"` // Lieux du concert
}

// Structure temporaire pour la désérialisation des données de l'API
var temp []struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
}

package groupie

type Website struct {
	Template string
}

type Artist struct {
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Date struct {
	ConcertDate []string `json:"dates"`
}

type Location struct {
	ConcertLocation []string `json:"locations"`
}

type Relation struct {
	DateLocation map[string][]string `json:"relation"`
}

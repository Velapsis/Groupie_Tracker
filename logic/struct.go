package groupie

type Website struct {
	Template string
}

type Artist struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relations    map[string][]string `json:"relations"`
}
type Relation struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}
type Concert struct {
	ConcertDate     []string `json:"dates"`
	ConcertLocation []string `json:"locations"`
}

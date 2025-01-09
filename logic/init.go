package groupie

func Init() {
	GetArtists()
	getRelations()
	

}

func InitPages() {
	MenuPage = "templates/menu.html"
	MainPage = "templates/game.html"
	ArtistPage = "templates/hardgame.html"
	ErrorPage = "templates/error.html"
}

func InitWeb() {
	InitPages()

}

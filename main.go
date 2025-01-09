package main

import (
	groupie "main/logic"
	server "main/templates"
)

func main() {
	groupie.InitWeb()
	server.CreateWebsite()

}

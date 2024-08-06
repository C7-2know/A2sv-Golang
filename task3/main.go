package main

import (
	"Library/controllers"
	"Library/services"
)


func main() {
	library := services.CreateLibrary()
	controllers.LibraryController(library)
}

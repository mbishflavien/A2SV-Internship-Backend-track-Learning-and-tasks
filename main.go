package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	// Initialize library service
	lib := services.NewLibrary()

	// Seed dummy member data for quick demonstration
	lib.AddMember(models.Member{ID: 1, Name: "Alice"})
	lib.AddMember(models.Member{ID: 2, Name: "Bob"})

	// Seed dummy book data
	lib.AddBook(models.Book{ID: 101, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Status: "Available"})
	lib.AddBook(models.Book{ID: 102, Title: "Clean Code", Author: "Robert C. Martin", Status: "Available"})

	// Start Controller
	controller := controllers.NewLibraryController(lib)
	controller.Run()
}
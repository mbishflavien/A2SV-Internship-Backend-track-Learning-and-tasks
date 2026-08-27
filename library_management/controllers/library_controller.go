package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	service services.LibraryManager
	reader  *bufio.Reader
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (c *LibraryController) Run() {
	for {
		fmt.Println("\n=== Library Management System ===")
		fmt.Println("1. Add a Book")
		fmt.Println("2. Remove a Book")
		fmt.Println("3. Borrow a Book")
		fmt.Println("4. Return a Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Member Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice (1-7): ")

		choiceStr := c.readLine()
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			c.addBook()
		case 2:
			c.removeBook()
		case 3:
			c.borrowBook()
		case 4:
			c.returnBook()
		case 5:
			c.listAvailableBooks()
		case 6:
			c.listBorrowedBooks()
		case 7:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select 1-7.")
		}
	}
}

func (c *LibraryController) readLine() string {
	input, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (c *LibraryController) readInt(prompt string) (int, error) {
	fmt.Print(prompt)
	valStr := c.readLine()
	return strconv.Atoi(valStr)
}

func (c *LibraryController) addBook() {
	id, err := c.readInt("Enter Book ID: ")
	if err != nil {
		fmt.Println("Invalid ID input.")
		return
	}

	fmt.Print("Enter Book Title: ")
	title := c.readLine()

	fmt.Print("Enter Book Author: ")
	author := c.readLine()

	c.service.AddBook(models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	})
	fmt.Println("Book added successfully!")
}

func (c *LibraryController) removeBook() {
	id, err := c.readInt("Enter Book ID to remove: ")
	if err != nil {
		fmt.Println("Invalid ID input.")
		return
	}

	if err := c.service.RemoveBook(id); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Book removed successfully!")
}

func (c *LibraryController) borrowBook() {
	bookID, err := c.readInt("Enter Book ID: ")
	if err != nil {
		fmt.Println("Invalid Book ID.")
		return
	}

	memberID, err := c.readInt("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}

	if err := c.service.BorrowBook(bookID, memberID); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Book borrowed successfully!")
}

func (c *LibraryController) returnBook() {
	bookID, err := c.readInt("Enter Book ID: ")
	if err != nil {
		fmt.Println("Invalid Book ID.")
		return
	}

	memberID, err := c.readInt("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}

	if err := c.service.ReturnBook(bookID, memberID); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Book returned successfully!")
}

func (c *LibraryController) listAvailableBooks() {
	books := c.service.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books in the library.")
		return
	}

	fmt.Println("\n--- Available Books ---")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func (c *LibraryController) listBorrowedBooks() {
	memberID, err := c.readInt("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}

	books, err := c.service.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("Member has no borrowed books.")
		return
	}

	fmt.Printf("\n--- Borrowed Books for Member %d ---\n", memberID)
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}
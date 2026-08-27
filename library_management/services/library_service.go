package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
	AddMember(member models.Member)
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	if book.Status == "" {
		book.Status = "Available"
	}
	l.Books[book.ID] = book
}

func (l *Library) AddMember(member models.Member) {
	l.Members[member.ID] = member
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	// Update Book status
	book.Status = "Borrowed"
	l.Books[bookID] = book

	// Add book to member's borrowed list
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	// Remove from member's borrowed slice
	found := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("book ID %d is not borrowed by member %s", bookID, member.Name)
	}

	// Update Book status back to Available
	book.Status = "Available"
	l.Books[bookID] = book
	l.Members[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exists := l.Members[memberID]
	if !exists {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}
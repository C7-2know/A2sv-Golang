package services

import (
	model "Library/models"
	"errors"
	"fmt"
	"strconv"
)

type LibraryManager interface {
	AddBook(book model.Book)
	RemoveBook(id int)
	BorrowBook(bookID int, memID int) error
	ReturnBook(bookID int, memID int) error
	ListAvailableBooks() []model.Book
	ListBorrowedBooks(memID int) []model.Book
}

type Library struct {
	books   map[int]model.Book
	members map[int]model.Member
}

func CreateLibrary() *Library {
	return &Library{books: make(map[int]model.Book), members: make(map[int]model.Member)}
}

func (l *Library) AddBook(book model.Book) {
	_, ok := l.books[book.ID]
	if ok {
		fmt.Println("Book already exists")
	} else {
		l.books[book.ID] = book
	}
}

func (l *Library) RemoveBook(id int) {
	for i, _ := range l.books {
		if i == id {
			delete(l.books, id)
			fmt.Println("You have removed the Book successfully")
			return
		}
	}
	fmt.Println("could not find the book")
}

func (l *Library) BorrowBook(bookID int, memID int) error {
	for i, book := range l.books {
		if i == bookID {
			if book.Status == "Borrowed" {
				return errors.New("Sorry! the Book is already borrowed")
			} else {
				for i , mem := range l.members {
					if i== memID {
						book.Status = "Borrowed"
						l.books[bookID] = book
						mem.BorrowedBooks = append(mem.BorrowedBooks, book)
						l.members[memID] = mem
						fmt.Println(l.books)
						fmt.Println("You have borrowed the book successfully")
						return nil
					}
				}
				newMem := model.Member{ID: memID, Name: "New"+strconv.Itoa(memID), BorrowedBooks: []model.Book{book}}
				l.members[memID] = newMem
				fmt.Println("you have borrowed the book and become a new member successfully")
				return nil
			}
		}
	}
	return errors.New("the Book does not found")
}

func (l *Library) ReturnBook(bookID int, memID int) error {
	for i, mem := range l.members {
		if i == memID {
			for i, book := range mem.BorrowedBooks {
				if book.ID == bookID {
					book.Status = "Available"
					mem.BorrowedBooks = append(mem.BorrowedBooks[:i], mem.BorrowedBooks[i+1:]...)
					l.members[memID] = mem
					l.books[bookID] = book
					fmt.Println("you have returned the book successfully")
					return nil
				}
			}
			return errors.New("We couldn't find the Book")
		}
	}
	return errors.New("The member does not exist")
}

func (l *Library) ListAvailableBooks() []model.Book {
	books:= []model.Book{}
	for _, book := range l.books {
		if book.Status == "Available" {
			books=append(books, book)
		}
	}
	return books
}

func (l *Library) ListBorrowedBooks(memID int) []model.Book {
	for _, mem := range l.members {
		if mem.ID == memID {
			return mem.BorrowedBooks
		}
	}
	return []model.Book{}
}

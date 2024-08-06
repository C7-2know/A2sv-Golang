package controllers

import (
	"Library/models"
	"Library/services"
	"fmt"
)

func LibraryController(libManager services.LibraryManager) {
	choice:=start()
	var id int
	var memID int
	for choice != 8 {
		switch choice {
		case 1:
			book := createBook()
			libManager.AddBook(book)
			fmt.Println("You have added the book successfully")
			choice = start()
			continue
			
		case 2:
			fmt.Println("Please enter the book id")
			fmt.Scanf("%d\n", &id)
			libManager.RemoveBook(id)
			choice = start()
			continue
		case 3:
			fmt.Println("Please enter the book id")
			fmt.Scanf("%d\n", &id)
			fmt.Println("Please enter the member id")
			fmt.Scanf("%d\n", &memID)
			err := libManager.BorrowBook(id, memID)
			if err != nil {
				fmt.Println(err)
			}
			choice = start()
			continue
		case 4:
			fmt.Println("Please enter the book id")
			fmt.Scanf("%d\n", &id)
			fmt.Println("Please enter the member id")
			fmt.Scanf("%d\n", &memID)
			err := libManager.ReturnBook(id, memID)
			if err != nil {
				fmt.Println(err)
			}
			choice = start()
			continue
		case 5:
			books := libManager.ListAvailableBooks()
			for _, book := range books {
				fmt.Println(book)
			}
			choice = start()
			continue
		case 6:
			fmt.Println("Please enter the member id")
			fmt.Scanf("%d\n", &memID)
			books := libManager.ListBorrowedBooks(memID)
			for _, book := range books {
				fmt.Println(book)
			}
			choice = start()
			continue
		}

	}
	fmt.Println("Goodbye")

}

func createBook() models.Book {
	var id int
	var author string
	var title string
	fmt.Println("Please enter the book id")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Please enter the book title")
	fmt.Scanf("%s\n", &title)
	fmt.Println(title,id)
	fmt.Println("Please enter the book author")
	fmt.Scanf("%s\n", &author)
	newBook := models.Book{ID: id,Title: title,Author: author, Status: "Available"}
	return newBook
}

func start() int{
	fmt.Println("Welcome to our library")
	fmt.Println("Please select an option")
	fmt.Println("1- Add a book")
	fmt.Println("2- Remove a book")
	fmt.Println("3- Borrow a book")
	fmt.Println("4- Return a book")
	fmt.Println("5- List all books")
	fmt.Println("6- List borrowed books")
	fmt.Println("8- Exit")
	var choice int
	fmt.Scanf("%d\n", &choice)
	return choice
}
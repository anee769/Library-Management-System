package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/anee769/Library-Management-System/Library"
)

func main() {
	var input int
	var err error
	flag := true
	lib, err := Library.NewLibrary()

	defer lib.Stop()

	fmt.Printf("\n\n\n\t\t\t\t\t\t\tWELCOME TO THE LIBRARY\n\n\n\n")
	fmt.Println("Codes for booktypes : ")
	fmt.Printf("1. eBook : %v\n2. AudioBook : %v\n3. HardBack : %v\n4. PaperBack : %v\n5. Encyclopedia : %v\n6. Magazine : %v\n7. Comic : %v\n\n\n", Library.EBook, Library.Audiobook, Library.Hardback, Library.Paperback, Library.Encyclopedia, Library.Magazine, Library.Comic)
	for flag {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Add Books to the Inventory")
		fmt.Println("2. Give books to borrow")
		fmt.Println("3. Receive back borrowed books")
		fmt.Println("4. Get Quantity of a book")
		fmt.Println("5. Get Author of a book")
		fmt.Println("6. Get BookType of a book")
		fmt.Println("7. Display the whole Inventory")
		fmt.Println("8. Display the Registered users")
		fmt.Println("9. Exit")
		fmt.Printf("Enter Input : ")
		fmt.Scan(&input)

		switch input {
		case 1:
			var (
				title, author, bookMedium string
				quantity                  int
				bookType                  Library.BookType
			)
			fmt.Println("Enter the Details of the Book : ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Println()
			fmt.Printf("Author : ")
			fmt.Scan(&author)
			fmt.Println()
			fmt.Printf("Quantity : ")
			fmt.Scan(&quantity)
			fmt.Println()
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Println()
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			fmt.Println()
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				err = lib.Pb.AddBook(title, author, quantity, bookType)
			} else {
				err = lib.Db.AddBook(title, author, quantity, bookType)
			}
			if err != nil {
				log.Fatalln(err)
			}

		case 2:
			var (
				title, author, bookMedium, borrower string
				quantity                            int
				bookType                            Library.BookType
			)
			fmt.Printf("Borrower Name : ")
			fmt.Scan(&borrower)
			fmt.Println("Enter the Details of the Book to be borrowed: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Println()
			fmt.Printf("Author : ")
			fmt.Scan(&author)
			fmt.Println()
			fmt.Printf("Quantity : ")
			fmt.Scan(&quantity)
			fmt.Println()
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Println()
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			fmt.Println()
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := Library.GenerateId(title, "Physical", bookType)
				fmt.Println(lib.Pb.Borrow(id, quantity, lib, borrower))
			} else {
				id := Library.GenerateId(title, "Digital", bookType)
				fmt.Println(lib.Db.Borrow(id, quantity, lib, borrower))
			}

		case 3:
			var (
				title, author, bookMedium, borrower string
				quantity                            int
				bookType                            Library.BookType
			)
			fmt.Printf("Borrower Name : ")
			fmt.Scan(&borrower)
			fmt.Println("Enter the Details of the Book to be returned: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Println()
			fmt.Printf("Author : ")
			fmt.Scan(&author)
			fmt.Println()
			fmt.Printf("Quantity : ")
			fmt.Scan(&quantity)
			fmt.Println()
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Println()
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			fmt.Println()
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := Library.GenerateId(title, "Physical", bookType)
				lib.Pb.Return(id, quantity, lib, borrower)
			} else {
				id := Library.GenerateId(title, "Digital", bookType)
				lib.Db.Return(id, quantity, lib, borrower)
			}

		case 4:
			var (
				title, bookMedium string
				bookType          Library.BookType
			)
			fmt.Println("Enter the Details of the Book whose quantity you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Println()
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Println()
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			fmt.Println()
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := Library.GenerateId(title, "Physical", bookType)
				fmt.Println(lib.Pb.GetQuantity(id))
			} else {
				id := Library.GenerateId(title, "Digital", bookType)
				fmt.Println(lib.Db.GetQuantity(id))
			}

		case 5:
			var (
				title, bookMedium string
				bookType          Library.BookType
			)
			fmt.Println("Enter the Details of the Book whose author you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Println()
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Println()
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			fmt.Println()
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := Library.GenerateId(title, "Physical", bookType)
				fmt.Println(lib.Pb.GetAuthor(id))
			} else {
				id := Library.GenerateId(title, "Digital", bookType)
				fmt.Println(lib.Db.GetAuthor(id))
			}

		case 6:
			var title string
			fmt.Println("Enter the Title of the Book whose bookType you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			lib.GetBookType(title)

		case 7:
			lib.ShowInventory()

		case 8:
			lib.ShowUsers()

		default:
			flag = false
		}
	}

	fmt.Printf("\n\n\n\n\t\t\t\t\t\t\tTHANK YOU!!\n\t\t\t\t\t\t\tVISIT AGAIN :)\n\n\n")

}

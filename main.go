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
	lib := Library.Library{}
	db := Library.DigitalBook{}
	pb := Library.PhysicalBook{}
	fmt.Printf("\t\t\t\t\t\t\tWELCOME TO THE LIBRARY\n\n\n\n")
	fmt.Println("Codes for booktypes : ")
	fmt.Printf("1. eBook : %v\n2. AudioBook : %v\n3. HardBack : %v\n4. PaperBack : %v\n5. Encyclopedia : %v\n6. Magazine : %v\n7. Comic : %v\n\n\n", eBook, Audiobook, Hardback, Paperback, Encyclopedia, Magazine, Comic)
	for flag {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Add Books to the Inventory")
		fmt.Println("2. Give books to borrow")
		fmt.Println("3. Receive back borrowed books")
		fmt.Println("4. Get Quantity of a book")
		fmt.Println("5. Get Author of a book")
		fmt.Println("6. Get BookType of a book")
		fmt.Println("7. Display the whole Inventory")
		fmt.Println("8. Exit")
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
			fmt.Printf("Author : ")
			fmt.Scan(&author)
			fmt.Printf("Quantity : ")
			fmt.Scan(&quantity)
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				err = pb.addBook(title, author, quantity, bookType)
			} else {
				err = db.addBook(title, author, quantity, bookType)
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
			fmt.Printf("Author : ")
			fmt.Scan(&author)
			fmt.Printf("Quantity : ")
			fmt.Scan(&quantity)
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := Library.generateId(title, "Physical", bookType)
				fmt.Println(pb.Borrow(id, quantity, lib, borrower))
			} else {
				id := Library.generateId(title, "Digital", bookType)
				fmt.Println(db.Borrow(id, quantity, lib, borrower))
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
			fmt.Printf("Author : ")
			fmt.Scan(&author)
			fmt.Printf("Quantity : ")
			fmt.Scan(&quantity)
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := Library.generateId(title, "Physical", bookType)
				pb.Return(id, quantity, lib, borrower)
			} else {
				id := Library.generateId(title, "Digital", bookType)
				db.Return(id, quantity, lib, borrower)
			}

		case 4:
			var (
				title, bookMedium string
				bookType          Library.BookType
			)
			fmt.Println("Enter the Details of the Book whose quantity you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := Library.generateId(title, "Physical", bookType)
				fmt.Println(pb.getQuantity(id))
			} else {
				id := Library.generateId(title, "Digital", bookType)
				fmt.Println(db.getQuantity(id))
			}

		case 5:
			var (
				title, bookMedium string
				bookType          Library.BookType
			)
			fmt.Println("Enter the Details of the Book whose author you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := Library.generateId(title, "Physical", bookType)
				fmt.Println(pb.getAuthor(id))
			} else {
				id := Library.generateId(title, "Digital", bookType)
				fmt.Println(db.getAuthor(id))
			}

		case 6:
			var title string
			fmt.Println("Enter the Title of the Book whose bookType you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			Library.GetBookType(title, db, pb)

		case 7:
			fmt.Println("Inventory : ")
			lib.showInventory(db, pb)

		default:
			flag = false
		}
	}

	fmt.Printf("\n\n\n\n\t\t\t\t\t\t\tTHANK YOU!!\n\t\t\t\t\t\t\tVISIT AGAIN :)\n\n\n")

}

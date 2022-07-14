package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

type BookType int64

const (
	eBook BookType = iota
	Audiobook
	Hardback
	Paperback
	Encyclopedia
	Magazine
	Comic
)

func generateId(title, bookMedium string, bookType BookType) string {
	h := sha1.New()
	h.Write([]byte(title + fmt.Sprintf("%v%v", bookType, bookMedium)))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

type Book interface {
	Details()
	Borrow()
	Return()
	addBook() error
	getAuthor() string
	getBookType() BookType
	getQuantity() int
}

type BookDetails struct {
	title      string
	author     string
	quantity   int
	bookType   BookType
	bookMedium string
}

type PhysicalBook struct {
	Inventory map[string]BookDetails
}

func (book PhysicalBook) Details() {
	for _, details := range book.Inventory {
		fmt.Printf("Title : %v\nAuthor : %v\nStock left : %v\nBookType : %v\nBookMedium : %v\n\n", details.title, details.author, details.quantity, details.bookType, details.bookMedium)
	}
}

func (book PhysicalBook) Borrow(id string, quantity int, lib Library, borrower string) string {

	if user, ok := lib.Users[borrower]; ok {
		if _, okIn := user[id]; okIn {
			return "Book already borrowed."
		} else if quantity > 1 {
			return "Cannot borrow more than 1 copy"
		}
	} else {
		lib.addMember(borrower)
	}

	if details, ok := book.Inventory[id]; ok {
		if details.quantity < quantity {
			return "Not Enough Quantity\n"
		} else {
			switch details.bookType {
			case Hardback:
			case Paperback:
			case Encyclopedia:
			case Magazine:
			case Comic:
			default:
				return "Invalid Booktype for Physical Book\n"
			}
			updatedQuantity := details.quantity - quantity
			book.Inventory[id] = BookDetails{book.Inventory[id].title, book.Inventory[id].author, updatedQuantity, details.bookType, "Physical"}
			lib.Users[borrower][id] = quantity
			return "Book Borrowed"
		}
	}
	return "This book is not present in stock\n"
}

func (book PhysicalBook) Return(id string, quantity int, lib Library, borrower string) {

	if details, ok := book.Inventory[id]; ok {
		details.quantity += quantity
		lib.Users[borrower][id] -= quantity
	}
	fmt.Println("Books Returned Successfully")
}

func (book PhysicalBook) addBook(title, author string, quantity int, bookType BookType) error {

	switch bookType {
	case Hardback:
	case Paperback:
	case Encyclopedia:
	case Magazine:
	case Comic:
	default:
		return fmt.Errorf("Invalid Booktype for Physical Book\n")
	}

	id := generateId(title, "Physical", bookType)
	if _, ok := book.Inventory[id]; ok {
		book.Inventory[id] = BookDetails{title, author, book.Inventory[id].quantity + quantity, bookType, "Physical"}
	} else {
		book.Inventory[id] = BookDetails{title, author, quantity, bookType, "Physical"}
	}
	return nil
}

func (book PhysicalBook) getAuthor(id string) string {
	if _, ok := book.Inventory[id]; ok {
		return fmt.Sprintf("Author of %v : %v", book.Inventory[id].title, book.Inventory[id].author)
	}
	return "This book is not present in Inventory"
}

func (book PhysicalBook) getBookType(title string) {
	flag := false
	for _, details := range book.Inventory {
		if details.title == title {
			fmt.Printf("Title : %v, BookType : %v, BookMedium : %v", details.title, details.bookType, details.bookMedium)
			flag = true
		}
	}
	if !flag {
		fmt.Println("This book is not present in Inventory")
	}
}

func (book PhysicalBook) getQuantity(id string) string {
	if _, ok := book.Inventory[id]; ok {
		return fmt.Sprintf("Stock left for %v : %v", book.Inventory[id].title, book.Inventory[id].quantity)
	}
	return "This book is not present in Inventory"
}

type DigitalBook struct {
	Inventory map[string]BookDetails
}

func (book DigitalBook) Details() {
	for _, details := range book.Inventory {
		fmt.Printf("Title : %v\nAuthor : %v\nStock left : %v\nBookType : %v\nBookMedium : %v\n\n", details.title, details.author, details.quantity, details.bookType, details.bookMedium)
	}
}

func (book DigitalBook) Borrow(id string, quantity int, lib Library, borrower string) string {

	if user, ok := lib.Users[borrower]; ok {
		if user[id]+quantity > 5 {
			return "Limit exceeded for borrowing books"
		}
	} else {
		lib.addMember(borrower)
	}

	if details, ok := book.Inventory[id]; ok {
		if details.quantity < quantity {
			return "Not Enough Quantity\n"
		} else {
			switch details.bookType {
			case Hardback:
			case Paperback:
			case Encyclopedia:
			case Magazine:
			case Comic:
			default:
				return "Invalid Booktype for Digital Book\n"
			}
			updatedQuantity := details.quantity - quantity
			book.Inventory[id] = BookDetails{book.Inventory[id].title, book.Inventory[id].author, updatedQuantity, details.bookType, "Physical"}
			lib.Users[borrower][id] += quantity
			return "Book Borrowed"
		}
	}
	return "This book is not present in stock\n"
}

func (book DigitalBook) Return(id string, quantity int, lib Library, borrower string) {

	if details, ok := book.Inventory[id]; ok {
		details.quantity += quantity
		lib.Users[borrower][id] -= quantity
	}
	fmt.Println("Books Returned Successfully")
}

func (book DigitalBook) addBook(title, author string, quantity int, bookType BookType) error {

	switch bookType {
	case eBook:
	case Audiobook:
	case Encyclopedia:
	case Magazine:
	case Comic:
	default:
		return fmt.Errorf("Invalid Booktype for Digital Book\n")
	}

	id := generateId(title, "Physical", bookType)
	if _, ok := book.Inventory[id]; ok {
		book.Inventory[id] = BookDetails{title, author, book.Inventory[id].quantity + quantity, bookType, "Digital"}
	} else {
		book.Inventory[id] = BookDetails{title, author, quantity, bookType, "Digital"}
	}
	return nil
}

func (book DigitalBook) getAuthor(id string) string {
	if _, ok := book.Inventory[id]; ok {
		return fmt.Sprintf("Author of %v : %v", book.Inventory[id].title, book.Inventory[id].author)
	}
	return "This book is not present in Inventory"
}

func (book DigitalBook) getBookType(title string) {
	flag := false
	for _, details := range book.Inventory {
		if details.title == title {
			fmt.Printf("Title : %v, BookType : %v, BookMedium : %v", details.title, details.bookType, details.bookMedium)
			flag = true
		}
	}
	if !flag {
		fmt.Println("This book is not present in Inventory")
	}
}

func (book DigitalBook) getQuantity(id string) string {
	if _, ok := book.Inventory[id]; ok {
		return fmt.Sprintf("Stock left for %v : %v", book.Inventory[id].title, book.Inventory[id].quantity)
	}
	return "This book is not present in Inventory"
}

type Library struct {
	Users map[string]map[string]int
}

func (lib Library) addMember(name string) {
	lib.Users[name][""] = 0
}

func GetBookType(title string, db DigitalBook, pb PhysicalBook) {
	db.getBookType(title)
	pb.getBookType(title)
}

func (lib Library) showInventory(db DigitalBook, pb PhysicalBook) {
	pb.Details()
	db.Details()
}

func main() {
	var input int
	var err error
	flag := true
	lib := Library{}
	db := DigitalBook{}
	pb := PhysicalBook{}
	fmt.Printf("\t\t\t\t\t\t\tWELCOME TO THE LIBRARY\n\n\n\n")
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
				bookType                  BookType
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
				bookType                            BookType
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
				id := generateId(title, "Physical", bookType)
				fmt.Println(pb.Borrow(id, quantity, lib, borrower))
			} else {
				id := generateId(title, "Digital", bookType)
				fmt.Println(db.Borrow(id, quantity, lib, borrower))
			}

		case 3:
			var (
				title, author, bookMedium, borrower string
				quantity                            int
				bookType                            BookType
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
				id := generateId(title, "Physical", bookType)
				pb.Return(id, quantity, lib, borrower)
			} else {
				id := generateId(title, "Digital", bookType)
				db.Return(id, quantity, lib, borrower)
			}

		case 4:
			var (
				title, bookMedium string
				bookType          BookType
			)
			fmt.Println("Enter the Details of the Book whose quantity you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := generateId(title, "Physical", bookType)
				fmt.Println(pb.getQuantity(id))
			} else {
				id := generateId(title, "Digital", bookType)
				fmt.Println(db.getQuantity(id))
			}

		case 5:
			var (
				title, bookMedium string
				bookType          BookType
			)
			fmt.Println("Enter the Details of the Book whose author you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			fmt.Printf("BookType : ")
			fmt.Scan(&bookType)
			fmt.Printf("BookMedium : ")
			fmt.Scan(&bookMedium)
			if bookMedium = strings.Trim(strings.ToLower(bookMedium), " "); bookMedium == "physical" {
				id := generateId(title, "Physical", bookType)
				fmt.Println(pb.getAuthor(id))
			} else {
				id := generateId(title, "Digital", bookType)
				fmt.Println(db.getAuthor(id))
			}

		case 6:
			var title string
			fmt.Println("Enter the Title of the Book whose bookType you want to know: ")
			fmt.Printf("Title : ")
			fmt.Scan(&title)
			GetBookType(title, db, pb)

		case 7:
			fmt.Println("Inventory : ")
			lib.showInventory(db, pb)

		default:
			flag = false
		}
	}

	fmt.Printf("\n\n\n\n\t\t\t\t\t\t\tTHANK YOU!!\n\t\t\t\t\t\t\tVISIT AGAIN :)\n\n\n")

}

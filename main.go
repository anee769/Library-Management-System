package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
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
		if user[id]+quantity > 5 {
			return "Sorry, you're exceeding your borrow limit."
		}
	} else {
		lib.addMember(borrower)
	}

	if details, ok := book.Inventory[id]; ok { //////needs to be improved
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
			if updatedQuantity == 0 {
				delete(book.Inventory, id)
			} else {
				book.Inventory[id] = BookDetails{book.Inventory[id].title, book.Inventory[id].author, updatedQuantity, details.bookType, "Physical"}
			}
		}
	}
	return "This book is not present in stock\n"
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
	book.Inventory[id] = BookDetails{title, author, quantity, bookType, "Physical"}
	return nil
}

func (book PhysicalBook) getAuthor(id string) string {
	if _, ok := book.Inventory[id]; ok {
		return fmt.Sprintf("Author of %v : %v", book.Inventory[id].title, book.Inventory[id].author)
	}
	return "This book is not present in Inventory"
}

func (book PhysicalBook) getBookType(id string) string {
	if _, ok := book.Inventory[id]; ok {
		return fmt.Sprintf("BookType of %v : %v", book.Inventory[id].title, book.Inventory[id].bookType) /////needs to be improved
	}
	return "This book is not present in Inventory"
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

	id := generateId(title, "Digital", bookType)
	book.Inventory[id] = BookDetails{title, author, quantity, bookType, "Digital"}
	return nil
}

func (book DigitalBook) getAuthor(id string) string {
	if _, ok := book.Inventory[id]; ok {
		return fmt.Sprintf("Author of %v : %v", book.Inventory[id].title, book.Inventory[id].author)
	}
	return "This book is not present in Inventory"
}

func (book DigitalBook) getBookType(id string) string { //////needs to be improved
	if _, ok := book.Inventory[id]; ok {
		return fmt.Sprintf("BookType of %v : %v", book.Inventory[id].title, book.Inventory[id].bookType)
	}
	return "This book is not present in Inventory"
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

func (lib Library) showInventory(b ...Book) {
	b[0].Details()
	b[1].Details()
}

func main() {
	fmt.Printf("\t\t\t\t\t\t\tWELCOME TO THE LIBRARY\n\n\n\n")

}

package Library

import "fmt"

type PhysicalBook struct {
	Inventory map[string]BookDetails
}

func (book PhysicalBook) Details() {
	for _, details := range book.Inventory {
		fmt.Printf("Title : %v\nAuthor : %v\nStock left : %v\nBookType : %v\nBookMedium : %v\n\n", details.title, details.author, details.quantity, details.bookType, details.bookMedium)
	}
}

func (book PhysicalBook) Borrow(id string, quantity int, lib Library, borrower string) string {

	for _, books := range lib.Users {
		if _, ok := books[id]; ok {
			return "Book already borrowed"
		}
	}

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
		if lib.Users[borrower][id] == 0 {
			delete(lib.Users[borrower], id)
		}
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

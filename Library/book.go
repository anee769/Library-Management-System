package Library

type BookType int64

const (
	EBook BookType = iota
	Audiobook
	Hardback
	Paperback
	Encyclopedia
	Magazine
	Comic
)

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

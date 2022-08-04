package Library

import (
	"fmt"

	badger "github.com/anee769/Library-Management-System/db"
)

type Library struct {
	Database *badger.Database

	Db    *DigitalBook
	Pb    *PhysicalBook
	Users map[string]map[string]int
}

func NewLibrary() (*Library, error) {
	lib := new(Library)

	if badger.Exists() {
		fmt.Println("exists")
		if err := lib.load(); err != nil {
			return nil, fmt.Errorf("failed to load existing lib: %w", err)
		}

	} else {
		fmt.Println("doesn't exist")
		if err := lib.init(); err != nil {
			return nil, fmt.Errorf("failed to initialize new lib: $%w", err)
		}
	}

	return lib, nil
}

func (lib *Library) load() (err error) {
	if lib.Database, err = badger.Open(); err != nil {
		return err
	}

	db, err := lib.Database.GetEntry([]byte("db"))
	if err != nil {
		return fmt.Errorf("lib dbook retrieve failed: %w", err)
	}

	pb, err := lib.Database.GetEntry([]byte("pb"))
	if err != nil {
		return fmt.Errorf("lib pbook retrieve failed: %w", err)
	}

	users, err := lib.Database.GetEntry([]byte("users"))
	if err != nil {
		return fmt.Errorf("user list retrieve failed: %w", err)
	}

	dobject, err := GobDecode(db, new(DigitalBook))
	if err != nil {
		return fmt.Errorf("error deserializing dbook : %w", err)
	}

	lib.Db = NewDigitalBook()
	lib.Db = dobject.(*DigitalBook)

	pobject, err := GobDecode(pb, new(PhysicalBook))
	if err != nil {
		return fmt.Errorf("error deserializing pbook : %w", err)
	}

	lib.Pb = NewPhysicalBook()
	lib.Pb = pobject.(*PhysicalBook)

	userobject, err := GobDecode(users, new(map[string]map[string]int))
	if err != nil {
		return fmt.Errorf("error deserializing pbook : %w", err)
	}

	lib.Users = make(map[string]map[string]int)
	lib.Users = *userobject.(*(map[string]map[string]int))

	return nil
}

func (lib *Library) init() (err error) {
	if lib.Database, err = badger.Open(); err != nil {
		return err
	}
	lib.Db = NewDigitalBook()
	lib.Pb = NewPhysicalBook()
	lib.Users = make(map[string]map[string]int)
	return nil
}

func (lib *Library) Stop() error {
	pb, err := GobEncode(lib.Pb)
	if err != nil {
		return fmt.Errorf("Encoding failed for Physical Book")
	}
	if err := lib.Database.SetEntry([]byte("pb"), pb); err != nil {
		return fmt.Errorf("Entry in database failed for pb %w", err)
	}

	db, err := GobEncode(lib.Db)
	if err != nil {
		return fmt.Errorf("Encoding failed for Digital Book")
	}
	if err := lib.Database.SetEntry([]byte("db"), db); err != nil {
		return fmt.Errorf("Entry in database failed for db %w", err)
	}

	users, err := GobEncode(lib.Users)
	if err != nil {
		return fmt.Errorf("Encoding failed for Users")
	}
	if err := lib.Database.SetEntry([]byte("users"), users); err != nil {
		return fmt.Errorf("Entry in database failed for users %w", err)
	}

	lib.Database.Close()
	return nil
}
func (lib *Library) AddMember(name string) {
	lib.Users[name] = make(map[string]int)
	fmt.Printf("\n%v added\n", name)
}

func (lib *Library) GetBookType(title string) {
	lib.Db.GetBookType(title)
	lib.Pb.GetBookType(title)
}

func (lib *Library) ShowInventory() {
	fmt.Printf("\n\n\n\t\t\tInventory\n\n\n")
	lib.Pb.Details()
	lib.Db.Details()
}

func (lib *Library) ShowUsers() {
	fmt.Printf("\n\n\n\t\t\tUser Inventory\n\n\n")
	for user := range lib.Users {
		fmt.Printf("%v : \n", user)
	}
}

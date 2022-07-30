package Library

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

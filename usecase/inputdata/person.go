package inputdata

type AddPerson struct {
	LoginID    string
	PersonName string
	Password   string
}

type UpdatePerson struct {
	LoginID    string
	PersonName string
	Password   string
}

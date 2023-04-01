package main

type user struct {
	Name, Address, Phone string
}

func (u *user) ChangeName(newName string) {
	u.Name = newName

}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress

}

type User interface {
	GhangeName(newName string)
	GhangeAddress(newAddress string)
}

func main() {

}

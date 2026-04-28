package minitask7

import "fmt"

type User struct {
	Name    string
	Address string
	Phone   string
}

func NewUser(name, address, phone string) *User {
	return &User{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (u *User) Print() {
	fmt.Printf("Nama : %s\nAddres : %s\nPhone : %s\n", u.Name, u.Address, u.Phone)
}

func (u *User) Greet() string {
	return ("Hai Salam Kenal " + u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}

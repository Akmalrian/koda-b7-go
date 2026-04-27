package minitask4

type User struct {
	Name      string
	Photo     string
	Email     string
	Age       int
	Phone     string
	Status    bool
	Education []Education
}

type Education struct {
	Nama    string
	Jurusan string
}

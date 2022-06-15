package poc

import "fmt"

func MethodsPoc() string {
	meTheUser := User{"Me", "me@me.com", true, 16}
	meTheUser.GetStatus()
	meTheUser.GetAge()

	return "---"
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) GetAge() {
	fmt.Println("Is user active:", u.Age)
}

package user

import (
	"fmt"
)

type User interface {
	Register()
	Unregister()
}

//Register variable number of Users
func Register(us ...User) {
	for _, currentUser := range us {
		currentUser.Register()
		fmt.Println("registered", us)
	}
}

//Unregister variable number of Users
func Unregister(us ...User) {
	for _, currentUser := range us {
		currentUser.Unregister()
	}
}

//TODO Return Password of User
func GetPassword(us User) string {
	return "placeholder"
}

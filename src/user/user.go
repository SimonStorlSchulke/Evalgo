package user

type User interface {
	Register()
	Unregister()
}

//Register variable number of Users
func Register(us ...User) {
	for _, currentUser := range us {
		currentUser.Register()
	}
}

//Unregister variable number of Users
func Unregister(us ...User) {
	for _, currentUser := range us {
		currentUser.Unregister()
	}
}

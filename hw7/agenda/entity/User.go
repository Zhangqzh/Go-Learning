package entity

type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

func GetPassword(user User) string {
	return user.Password
}

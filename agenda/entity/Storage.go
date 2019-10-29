package entity

import (
	"fmt"
	"regexp"
)

var Users []User

func init() {

	user := ReadFromFile()
	for i := 0; i < len(user); i++ {
		Users = append(Users, user[i])
	}

}

func IsEmailAvailable(email string) bool {
	var flag bool
	flag, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email)
	if flag == false {
		fmt.Println("Please use an available email!")
	}
	return flag
}

func IsCellphoneAvailable(phone string) bool {
	var flag bool
	flag, _ = regexp.MatchString("^1[0-9]{10}$", phone)
	if flag == false {
		fmt.Println("Please use an available phone number!")
	}
	return flag
}

func isPasswordAvailable(password string) bool {
	var flag bool
	flag = true
	if len(password) < 6 {
		fmt.Println("Please use a password longer than 6! ")
		flag = false
	}
	return flag
}

func isUserExists(name string) bool {

	flag, _ := QueryUser(name)
	if flag {
		fmt.Println("This username already exits, please use another username like name123")
	}
	return flag
}

func UserRegister(name string, password string, email string, phone string) bool {
	var user User
	if isPasswordAvailable(password) && IsEmailAvailable(email) && IsCellphoneAvailable(phone) && !isUserExists(name) {
		user.Name = name
		user.Password = password
		user.Email = email
		user.Phone = phone
		Users = append(Users, user)
		WriteToFile(Users)
		fmt.Println("Register successfully!")
		return true
	} else {
		return false
	}
}

func QueryUser(name string) (bool, User) {

	for i := 0; i < len(Users); i++ {
		if Users[i].Name == name {
			return true, Users[i]
		}
	}
	return false, User{"none", "none", "none", "none"}
}

package service

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/agenda/entity"
)


var log_file *os.File



func init() {
	logFile, err := os.OpenFile("files/agenda.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	log_file = logFile
	if err != nil {
		log.Fatalln("open file error !")
	}

	

}

func RegisterUser(name string, password string, email string, phone string) {

	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	re := entity.UserRegister(name, password, email, phone)
	if re {
		debugLog.Println(name, " register successfully!")
	} else {
		debugLog.Println(name, " register failed!")
	}

	defer log_file.Close()
}

func Log_in(name string, password string) {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	flag, logUser := entity.QueryUser(name)
	//fmt.Println(name)
	if flag == true {
		if entity.GetPassword(logUser) != password {
			debugLog.Println(name, " log in failed!")
			fmt.Println("Password is wrong!")
		} else {
			debugLog.Println(name, " log in successfully!")
			fmt.Println("Log in successfully!\n")
		}
	} else {
		debugLog.Println(name, " log in failed!")
		fmt.Println("No user in list!")
	}
	
	defer log_file.Close()
}

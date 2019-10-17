package entity

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func JsonDecode(js []byte) User {
	var jm User
	err := json.Unmarshal(js, &jm)
	if err != nil {
		fmt.Println("error2")
	}
	return jm
}

func JsonEncode(m User) []byte {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error1")
		os.Exit(1)
	}
	return data
}

func ReadFromFile() []User {
	var user []User
	f, err := os.Open("files/User.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		user = append(user, JsonDecode([]byte(line)))
	}
	return user
}

func WriteToFile(user []User) {
	file, err := os.OpenFile("files/User.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate("files/User.txt", 0)
	if err != nil {
		fmt.Println("open file failed.", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(user); i++ {
		file.WriteString(string(JsonEncode(user[i])[:]))
		file.WriteString("\n")
	}
}



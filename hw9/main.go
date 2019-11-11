package main

import (
	"os"

	flag "github.com/spf13/pflag"
	"github.com/user/hw9/service"
)

//默认监听端口
const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}
	//命令行输入监听端口，绑定、解析端口
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	//启动服务器
	service.NewServer(port)
}

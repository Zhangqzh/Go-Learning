package service

import (
	"github.com/codegangsta/martini"
)

func NewServer(port string) {
	r := martini.Classic()
	//提交请求的处理
	r.Get("/", func(params martini.Params) string {
		return "Hello world"
	})

	r.RunOnAddr(":" + port)
}

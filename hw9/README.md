
# 项目地址：
[GoOnline]()  
[Github](https://github.com/Zhangqzh/Go-Learning/tree/master/hw9)  
# 项目概述 
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。
## 任务目标
1. 熟悉 go 服务器工作原理  
2. 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。  
3. 使用 curl 工具访问 web 程序  
4. 对 web 执行压力测试  
## 相关知识
课件：http://blog.csdn.net/pmlpml/article/details/78404838
# 任务要求
## 基本要求
1. 编程 web 服务程序 类似 cloudgo 应用。
    - 要求有详细的注释
    - 是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
2. 使用 curl 测试，将测试结果写入 README.md
3. 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

# 实验过程
## 代码及注释
### servce.go
```go
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
```
### main.go
```go
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
```
## 使用框架说明
### net/http
Go 语言中处理 HTTP 请求主要跟两个东西相关：ServeMux 和 Handler。

ServrMux 本质上是一个 HTTP 请求路由器（或者叫多路复用器，Multiplexor）。它把收到的请求与一组预先定义的 URL 路径列表做对比，然后在匹配到路径的时候调用关联的处理器（Handler）。

处理器（Handler）负责输出HTTP响应的头和正文。任何满足了http.Handler接口的对象都可作为一个处理器。通俗的说，对象只要有个如下签名的ServeHTTP方法即可：
```go
ServeHTTP(http.ResponseWriter, *http.Request)
```
包http提供HTTP客户端和服务器实现  
Get，Head，Post和PostForm发出HTTP（或HTTPS）请求  完成后，客户端必须关闭相应主体

### Martini
Martini是一个强大为了编写模块化Web应用而生的GO语言框架.  
#### 功能列表：
- 使用极其简单.
- 无侵入式的设计.
- 很好的与其他的Go语言包协同使用.
- 超赞的路径匹配和路由.
- 模块化的设计 - 容易插入功能件，也容易将其拔出来.
- 已有很多的中间件可以直接使用.
- 框架内已拥有很好的开箱即用的功能支持.
- 完全兼容[http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)接口.  
参考链接：[Martini](https://github.com/go-martini/martini/blob/master/translations/README_zh_cn.md)  
martini 是新锐的框架，只是一个微型框架，只带有简单的核心，路由功能和依赖注入容器inject。但目前我们也不需要自己写依赖什么的，也不用和数据库结合使用所以还是选择简单的。
## 运行结果
### 运行测试
```bash
go run main.go -p 9000
```
![](https://www.z4a.net/images/2019/11/11/TIM20191111152326.png)  
监听端口为9000，在浏览器中输入`http://localhost:9000`可以看到hello world  

![](https://www.z4a.net/images/2019/11/11/TIM20191111151912.png)  
### curl测试
`curl -v http://localhohst:9000`
![](https://www.z4a.net/images/2019/11/11/TIM20191111152017.png)
### ab测试
安装Apache web 压力测试程序  
`yum -y install httpd-tools`
`ab -n 1000 -c 100 http://localhost:9000`
![](https://www.z4a.net/images/2019/11/11/TIM20191111152046.png)
命令行参数：
- -n：执行的请求数量
- -c: 并发请求个数
- -t：测试所进行的最大秒数
- -p：包含了需要POST的数据的文件
- -T：POST数据所使用的Content-type头信息
- -k：启用HTTP KeepAlive功能，即在一个HTTP会话中执行多个请求，默认时，不启用KeepAlive功能  

结果参数：
- Server Software: 服务器软件版本  
- Server Hostname: 请求的URL  
- Server Port: 请求的端口号  
- Document Path: 请求的服务器的路径  
- Document Length: 页面长度 单位是字节  
- Concurrency Level: 并发数  
- Time taken for tests: 一共使用了的时间  
- Complete requests: 总共请求的次数  
- Failed requests: 失败的请求次数  
- Total transferred: 总共传输的字节数 http头信息  
- HTML transferred: 实际页面传递的字节数  
- Requests per second: 每秒多少个请求  
- Time per request: 平均每个用户等待多长时间  
- Time per request: 服务器平均用多长时间处理  
- Transfer rate: 传输速率  
- Connection Times: 传输时间统计  
- Percentage of the requests served within a certain time: 确定时间内服务请求占总数的百分比  


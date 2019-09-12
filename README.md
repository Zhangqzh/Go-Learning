# Go-Learning

## **Golang的配置和Git的连接**
[配置过程和遇到的问题](https://zhangqzh.github.io/2019/09/10/Linux%E4%B8%8B%E9%85%8D%E7%BD%AEGolang%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E5%92%8CGit%E8%BF%9C%E7%A8%8B%E4%BB%93%E5%BA%93/)写在博客里  
## **编写Go程序**

### 工作空间
> 工作空间是一个存放管理Go代码的目录，其中包含三个子目录：
src目录包含Go的源文件，每个目录都代表了一个源码包  
pkg目录包含包对象，存放编译后的包和依赖包  
bin目录包含可执行文件，存放可执行对象  
go工具用于构建源码包，并将其生成的二进制文件安装到pkg和bin目录中  
src子目录常会包含多种版本控制的代码仓库，如Git或Mercurial，以此来跟踪一个或多个源码包的开发

如下：
```
bin/
	streak                         # 可执行命令
	todo                           # 可执行命令
pkg/
	linux_amd64/
		code.google.com/p/goauth2/
			oauth.a                # 包对象
		github.com/nf/todo/
			task.a                 # 包对象
src/
	code.google.com/p/goauth2/
		.hg/                       # mercurial 代码库元数据
		oauth/
			oauth.go               # 包源码
			oauth_test.go          # 测试源码
	github.com/nf/
		streak/
		.git/                      # git 代码库元数据
			oauth.go               # 命令源码
			streak.go              # 命令源码
		todo/
		.git/                      # git 代码库元数据
			task/
				task.go            # 包源码
			todo.go                # 命令源码
```
  
### 编写并运行hello.go

- 要编译并运行简单的程序，首先要选择包路径（我们在这里使用 github.com/user/hello），并在你的工作空间内创建相应的包目录:  

```
$ mkdir $GOPATH/src/github.com/user/hello
```
- 在该目录中创建名为hello.go的文件，其内容为以下代码：
``` go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```
- 用go工具构建并安装此程序
```
$ go install github.com/user/hello
```
注意，你可以在系统的任何地方运行此命令。go 工具会根据 GOPATH 指定的工作空间，在 github.com/user/hello 包内查找源码.  
若在从包目录中运行 go install，也可以省略包路径：

```
$ cd $GOPATH/src/github.com/user/hello
$ go install
```

- 输入完整路径运行
```
$ $GOPATH/bin/hello
Hello, world.
```

![](http://pxcpbo9xv.bkt.clouddn.com/01.png)
## **将代码推送到远程仓库**
### git的配置以及两种连接方式和遇到的问题已经写在博客里了

```
$ cd $GOPATH/src/github.com/user/hello
$ git init
$ git add hello.go
$ git commit -m "initial commit"
$ git remote add origin https://github.com/user-name/repo-name.git
$ git remote -v
$ git push -u origin master
```
### 实验结果
![](http://pxcpbo9xv.bkt.clouddn.com/02.png)
![](http://pxcpbo9xv.bkt.clouddn.com/03.png)

这里我第一次push的时候报错了，好像是因为库里本来就有了一个hello.go，但是我去看了并没有，，很迷惑  
更改指令为
```
git push -f origin master
```
**强推**，即利用强覆盖方式用你本地的代码替代git仓库内的内容  
参见[【GIT】常用GIT知识点](https://github.com/zhongxia245/blog/issues/14)


## **你的第一个库**
### 编写一个库，并让`hello`程序来使用它
- 创建包目录
```
mkdir $GOPATH/src/github.com/user/stringutil
```
- 在该目录下创建名为`reverse.go`的文件  
代码如下：  
``` go
// stringutil 包含有用于处理字符串的工具函数。
package stringutil

// Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```
- 使用`go build`命令来测试该包的编译
```
$ go build github.com/user/stringutil
```
这不会产生输出文件，想要输出的话，必须使用`go install`命令，它会将包的对象放到工作空间`pkg`目录中  
![](http://pxcpbo9xv.bkt.clouddn.com/04.png)   
- 确认`stringutil`包构建完毕后，修改原来的`hello.go`文件
```go
package main

import (
	"fmt"

	"github.com/user/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
```
- 无论是安装包还是二进制文件，`go`工具都会安装它所依赖的任何东西，因此当我们通过
```
go install github.com/user/hello
```
来安装`hello`程序时,`stringutil`包也会被自动安装  

- 运行结果  
![](http://pxcpbo9xv.bkt.clouddn.com/05.png)
- 做完这些之后，工作空间变为
```
bin/
	hello                 # 可执行命令
pkg/
	linux_amd64/          # 这里会反映出你的操作系统和架构
		github.com/user/
			stringutil.a  # 包对象
src/
	github.com/user/
		hello/
			hello.go      # 命令源码
		stringutil/
			reverse.go       # 包源码
```
## **测试**
通过创建文件`reverse_test.go`来为`stringutil`添加测试,文件同样应该在`stringutil`目录下  
代码如下：
```go
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```
运行该测试
![](http://pxcpbo9xv.bkt.clouddn.com/06.png)  
## **远程包**
`go`工具可从远程代码库自动获取包  
若你在包的导入路径中包含了代码仓库的URL，`go get`就会自动获取，构建，并安装它
```
$ go get github.com/golang/example/hello
$ $GOPATH/bin/hello
Hello, Go examples!
```
若指定的包不在工作空间中，`go get`就会将它放到`GOPATH`指定的第一个工作空间内。（若该包已存在，`go get`就会跳过远程获取，其行为与`go install`相同）    
执行结果:
![](http://pxcpbo9xv.bkt.clouddn.com/07.png)  

工作目录：
```
bin/
	hello                 # 可执行命令
pkg/
	linux_amd64/
		code.google.com/p/go.example/
			stringutil.a     # 包对象
		github.com/user/
			stringutil.a     # 包对象
src/
	code.google.com/p/go.example/
		hello/
			hello.go      # 命令源码
		stringutil/
			reverse.go       # 包源码
			reverse_test.go  # 测试源码
	github.com/user/
		hello/
			hello.go      # 命令源码
		stringutil/
			reverse.go       # 包源码
			reverse_test.go  # 测试源码
```

**此文件是在虚拟机上跑用的，所以import的地址和Go_online不太一样**  
目录：
- cobra用法解释
- Linux虚拟机里开发使用教程
- Go-online使用说明
- 部分代码解释
- 实验结果

# 准备工作
## 1. 概述
命令行实用程序并不是都象 cat、more、grep 是简单命令。go 项目管理程序，类似 java 项目管理 maven、Nodejs 项目管理程序 npm、git 命令行客户端、 docker 与 kubernetes 容器管理工具等等都是采用了较复杂的命令行。即一个实用程序同时支持多个子命令，每个子命令有各自独立的参数，命令之间可能存在共享的代码或逻辑，同时随着产品的发展，这些命令可能发生功能变化、添加新命令等。因此，符合 OCP 原则 的设计是至关重要的编程需求。  
## 2. JSON序列化与反序列化
参考：[JSON and Go](https://blog.go-zh.org/json-and-go)
json 包是内置支持的，文档位置：https://go-zh.org/pkg/encoding/json/
## 3. 复杂命令行的处理
不要轻易“发明轮子”。为了实现POSIX/GNU-风格参数处理，-flags，包括命令完成等支持，程序员们开发了无数第三方包，这些包可以在[godoc](https://godoc.org/)找到。  
- pflag 包： https://godoc.org/github.com/spf13/pflag
- cobra 包： https://godoc.org/github.com/spf13/cobra
- goptions 包： https://godoc.org/github.com/voxelbrain/goptions
- docker command 包：https://godoc.org/github.com/docker/cli/cli/command
- ……
[go dead project](https://www.xuebuyuan.com/1588520.html)非常有用
这里我们选择cobar这个工具  


## 4. 安装并使用cobra
（在虚拟机里，我做的时候还不知道必须得在Go-online上提交）
### cobra安装
在`$GOPATH/src/golang.org/x`目录下用`git clone`下载`sys`和`text`项目，然后使用`go get -v github.com/spf13/cobra/cobra`  
下载成功后会在`$GOPATH/bin`目录下出现cobra可执行程序  
执行cobra，如图所示，即为成功安装  
![](https://www.z4a.net/images/2019/10/17/57222d9735dcbee9f.png)  

### cobra init
参考[官方文档](https://github.com/spf13/cobra#overview)  
生成agenda项目  
这里之前老是不可，然后点进去官方文档里的[Using Cobra Generator](https://github.com/spf13/cobra/blob/master/cobra/README.md),里面有详细的解释说明  
`./cobra init --pkg-name github.com/spf13/agenda`   
cobra init [app]命令创建初始应用程序代码，用正确的结构填充程序，并自动将LICENSE应用到程序中。  
在cobra应用程序中，通常main.go是暴露的文件，用它来初始化cobra，仅仅调用executecmd包的功能


### cobra add
添加agenda工具命令（只完成注册和登陆两个功能）
`./cobra add register`  
`./cobra add login`  
```go
package main

import "github.com/spf13/agenda/cmd"

func main() {
  cmd.Execute()
}
```
### 项目结构
```
agenda
  |─ cmd
  |  |─ register.go
  |  |─ login.go
  |  └─ root.go
  |─ entity
  |  |─ fileIO.go
  |  |─ Storage.go
  |  |─ User.go
  |─ files
  |  |─ User.txt
  |  └─ agenda.log
  |─ service
  |  └─ service.go 
  |─ LICENSE
  |─ agenda.io
  └─ main.go
```
（这里的agenda.o是后来手工移到这个文件夹的）  
项目结构就是按照去年实训的时候设计的
- User.go：存储User数据结构和返回方法
- fileIO.go：将User信息的文档形式和程序数据结构互相转化
- Storage.go：注册数据结构
- Service.go：被login.go和root.go调用，实现函数基础功能
- register.go：用户注册添加的app
- login.go：用户登陆添加的app  


## 5. Go-online使用
感谢群里喜欢提问的同学和回答问题的助教，让我一个小白通过看群消息知道了路径和依赖问题  
> go只在$GOPATH/src目录开始找依赖包而gopath可以有多个，多个的话就用冒号分割，这时候pwd命令看当前目录，就在第一项的src/的子目录，按顺序查找，而go文档的话是没有提到相对路径这种说法的，所以所有包的查找都是用这种方式，比如有一个包在GOPATH/src/a/b/x目录，那么引用就是import a/b/x，不用修改GOPATH，只用在import的时候写好项目路径就可以了，通过pwd查看项目路径就可以了  

以上都是群里的话直接搬过来的，助教真的解释的很详细，加上好像有好心的大佬已经解决了我上面在虚拟机那里说的依赖问题，就只用`go get -v github.com/spf13/cobra/cobra`项目就可以跑起来了
# Agenda程序开发
## 选项（Flag）
实际命令都有选项，分为持久和本地，持久例如kubectl的-n可以用在很多个二集命令下，本地命令选项则不会被继承到子命令。  

- type也有Slice，Count Duration,IP,IPMask,IPNet之类的类型,Slice类型可以多个传入，直接获取就是一个切片，例如–master ip1 –master ip2  
- 类似--force这样的开关型选项，实际上用Bool类型即可，默认值设置为false，单独给选项不带值就是true，也可以手动传入false或者true  
- MarkDeprecated告诉用户放弃这个标注位，应该使用新标志位MarkShorthandDeprecated是只放弃短的，长标志位依然可用。MarkHidden隐藏标志位
- MarkFlagRequired(“region”)表示region是必须的选项，不设置下选项都是可选的

## 日志服务
使用[log包](https://go-zh.org/pkg/log/)记录日志  
## 正则表达式（识别合法邮箱和号码）
导入regexp包，参考博客[Golang-regexp包](https://www.cnblogs.com/golove/p/3270918.html)  

## 部分代码
### cmd/root.go
```go
/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "agenda",
  Short: "Agenda is an application that manage the information of workers and meetings",
  Long: `This one just finish the function of user registering and logging in`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.agenda.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".agenda" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".agenda")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

```

### register.go
```go
/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/agenda/service"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register -n [username] -p [password] -e [email] -t [phone]",
	Short: "Register a new user",
	Long:  `Input command model like: register -n User -p 123456(longer than 6) -e 123@qq.com -t 1**********`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		u_name, _ := cmd.Flags().GetString("name")
		u_password, _ := cmd.Flags().GetString("password")
		u_email, _ := cmd.Flags().GetString("email")
		u_phone, _ := cmd.Flags().GetString("phone")

		service.RegisterUser(u_name, u_password, u_email, u_phone)

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	registerCmd.Flags().StringP("name", "n", "", "user name")
	registerCmd.Flags().StringP("password", "p", "", "user password")
	registerCmd.Flags().StringP("email", "e", "", "user email")
	registerCmd.Flags().StringP("phone", "t", "", "user phone")

}

```
rootCmd为init的root.go定义的结构体，rootCmd.AddCommand(appCmd)这里字面意思可以得知command这个结构体生成对应的命令格式，可以用上一层次的命令方法AddCommand添加一个下一级别的命令  

### login.go 
```go
/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/agenda/service"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "log_in -n [username] -p [password]",
	Short: "log in",
	Long:  `Input command mode like : log_in -n User -p 123456`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		u_name, _ := cmd.Flags().GetString("name")
		u_password, _ := cmd.Flags().GetString("password")
		service.Log_in(u_name, u_password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("name", "n", "", "user name")
	loginCmd.Flags().StringP("password", "p", "", "user password")

}

```  

## 实验结果
`./agenda`  
虚拟机运行：
![](https://www.z4a.net/images/2019/10/17/137e614d97e11aeafb.png)  
Go-online运行:  
![](https://www.z4a.net/images/2019/10/26/_20191026204938.png)      
`./agenda register -h`  
虚拟机运行：  
![](https://www.z4a.net/images/2019/10/17/145f737d1a5a3113b4.png)  
Go-online运行:  
![](https://www.z4a.net/images/2019/10/26/TIM20191026205218.png)    
指令提示  
虚拟机运行： 
![](https://www.z4a.net/images/2019/10/17/7ac1cdf713a58b59c.png)   
Go-online运行:  
![](https://www.z4a.net/images/2019/10/26/TIM20191026205318.png)     
`./agenda register -n -p -e -t`   
 虚拟机运行：  
![](https://www.z4a.net/images/2019/10/17/88ae6032f92793486.png)  
Go-online运行:  
![](https://www.z4a.net/images/2019/10/26/TIM20191026205520.png)    
`./agenda log_in -h`  
虚拟机运行：  
![](https://www.z4a.net/images/2019/10/17/9426b8502dab98fd8.png)   
 Go-online运行:   
 ![](https://www.z4a.net/images/2019/10/26/TIM20191026205619.png)    
`./agenda log_in -n -p`  
密码正确   
虚拟机运行：  
![](https://www.z4a.net/images/2019/10/17/105b4270fae0e29005.png)  
Go-online运行:  
![](https://www.z4a.net/images/2019/10/26/TIM20191026205717.png)  
密码错误  
虚拟机运行： 
![](https://www.z4a.net/images/2019/10/17/11bf4e2cee95d62da3.png)  
Go-online运行:    

User.txt  
![](https://www.z4a.net/images/2019/10/17/15604aa020a8098797.png)  
![](https://www.z4a.net/images/2019/10/26/_20191026210936.png)  
Agenda.log   
![](https://www.z4a.net/images/2019/10/17/12dfd5a6ccf2ba3c36.png)  
![](https://www.z4a.net/images/2019/10/26/_20191026212212.png)  

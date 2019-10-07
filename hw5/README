## 1. 概述
CLI（Command Line Interface）实用程序是Linux下应用开发的基础。正确的编写命令行程序让应用与操作系统融为一体，通过shell或script使得应用获得最大的灵活性与开发效率。Linux提供了cat、ls、copy等命令与操作系统交互；go语言提供一组实用程序完成从编码、编译、库管理、产品发布全过程支持；容器服务如docker、k8s提供了大量实用程序支撑云服务的开发、部署、监控、访问等管理任务；git、npm等都是大家比较熟悉的工具。尽管操作系统与应用系统服务可视化、图形化，但在开发领域，CLI在编程、调试、运维、管理中提供了图形化程序不可替代的灵活性与效率。
主要内容实现在第一个链接
[开发Linux命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)

## 2. selpg
[selpg.c源码](https://www.csdn.net/link?target_url=https%3A%2F%2Fwww.ibm.com%2Fdeveloperworks%2Fcn%2Flinux%2Fshell%2Fclutil%2Fselpg.c&id=82927996&token=c3b5a946f71045b9199bd35bc395aaf8)  

### flag
[Golang之使用Flag和Pflag](https://www.csdn.net/link?target_url=https%3A%2F%2Fo-my-chenjian.com%2F2017%2F09%2F20%2FUsing-Flag-And-Pflag-With-Golang%2F&id=82927996&token=7d81520b8ce986586aa06b4a09fcbf61)  
[Package pflag](https://www.csdn.net/link?target_url=https%3A%2F%2Fgodoc.org%2Fgithub.com%2Fspf13%2Fpflag%23Parse&id=82927996&token=c646f066b4e3534a91474203448abd9d)  
相关代码：
```go

	pflag.IntVarP(&(args.startPage), "startPage", "s", -1, "Define startPage")
	pflag.IntVarP(&(args.endPage), "endPage", "e", -1, "Define endPage")
	pflag.IntVarP(&(args.pageLen), "pageLength", "l", 72, "Define pageLength")
	pflag.StringVarP(&(args.printDest), "printDest", "d", "", "Define printDest")
    pflag.BoolVarP(&(args.pageType), "pageType", "f", false, "Define pageType")
```  
### 检查变量和参数名
```go
func checkArgs(args *selpgArgs) {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "\n[Error]The arguments are not enough!\n")
		pflag.Usage()
		os.Exit(1)
	} else if (args.startPage == -1) || (args.endPage == -1) {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage can't be empty! Please check your command!\n")
		pflag.Usage()
		os.Exit(2)
	} else if (args.startPage <= 0) || (args.endPage <= 0) {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage can't be negative! Please check your command!\n")
		pflag.Usage()
		os.Exit(3)
	} else if args.startPage > args.endPage {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage can't be bigger than the endPage! Please check your command!\n")
		pflag.Usage()
		os.Exit(4)
	} else if (args.pageType == true) && (args.pageLen != 72) {
		fmt.Fprintf(os.Stderr, "\n[Error]The command -l and -f are exclusive, you can't use them together!\n")
		pflag.Usage()
		os.Exit(5)
	} else if args.pageLen <= 0 {
		fmt.Fprintf(os.Stderr, "\n[Error]The pageLen can't be less than 1 ! Please check your command!\n")
		pflag.Usage()
		os.Exit(6)
	} else {
		pageType := "page length."
		if args.pageType == true {
			pageType = "The end sign /f."
		}
		fmt.Printf("\n[ArgsStart]\n")
		fmt.Printf("startPage: %d\nendPage: %d\ninputFile: %s\npageLength: %d\npageType: %s\nprintDestation: %s\n[ArgsEnd]", args.startPage, args.endPage, args.inFileName, args.pageLen, pageType, args.printDest)
	}

}
```
### OS
Command
参考资料
[golang中os/exec包的用法](https://blog.csdn.net/chenbaoke/article/details/42556949)  
```go
func excuteCMD(args *selpgArgs) {
	var fin *os.File
	if args.inFileName == "" {
		fin = os.Stdin
	} else {
		checkFileAccess(args.inFileName)
		var err error
		fin, err = os.Open(args.inFileName)
		checkError(err, "File input")
	}

	if len(args.printDest) == 0 {
		output2Des(os.Stdout, fin, args.startPage, args.endPage, args.pageLen, args.pageType)
	} else {
		output2Des(cmdExec(args.printDest), fin, args.startPage, args.endPage, args.pageLen, args.pageType)
	}
}
```
通过调用exec.Command()执行命令，返回一个cmd的结构体指针，cmd.StdinPipe()返回一个连接到command标准输入的管道，cmd.StoutPipe()返回一个连接到command标准输出的管道Pipe  
使用fout通过command管道向printDest文件写入数据  
cmd.Start()使用某个命令开始执行
## 3. 使用测试
`./hw51 -s1 -e1 input_file.txt`  
!()[1]
`./hw51 -s1 -e1 <input_file.txt`  
!()[2]
`./hw51 -s1 -e2 input_file.txt >output_file`  
!()[3]
`./hw51 -s1 -e4 input_file.txt 2>error_file`  
!()[4]
`./hw51 -s1 -e3 input_file.txt >output_file 2>error_file`  
!()[5]
`./hw51 -s1 -e2 -f input_file.txt`  
!()[6]  
测试文件包含一个换页符
!()[7]

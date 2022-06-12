go语言的安装： 下载地址：
https://studygolang.com/dl/golang/go1.18.3.windows-amd64.msi
go run *.go 运行文件 go build *.go 编译文件

程序入口 包名一定是要： package main 方法必须是： func main()
main 函数不支持返回值
无法使用直接传入命令行参数
os.Args[0] 是应用程序入口的路径

可以测试的方法：
文件名： 已_test结尾
method名已Test开头
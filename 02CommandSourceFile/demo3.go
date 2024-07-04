// 自定义命令源码文件的参数使用说明
package main

import (
	"flag"
	"fmt"
	"os"
)

var name2 string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	// 函数flag.StringVar接受 4 个参数
	// 第 1 个参数是用于存储该命令参数值的地址
	// 第 2 个参数是为了指定该命令参数的名称
	// 第 3 个参数是为了指定在未追加该命令参数时的默认值
	// 第 4 个参数是该命令参数的简短说明
	flag.StringVar(&name2, "name", "everyone", "The greeting object.")
}
func main() {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	flag.Parse()
	fmt.Printf("Hello, %s\n", name2)
}

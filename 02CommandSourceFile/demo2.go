package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	// 函数flag.StringVar接受 4 个参数
	// 第 1 个参数是用于存储该命令参数值的地址
	// 第 2 个参数是为了指定该命令参数的名称
	// 第 3 个参数是为了指定在未追加该命令参数时的默认值
	// 第 4 个参数是该命令参数的简短说明
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
func main() {
	flag.Parse()
	fmt.Printf("Hello, %s\n", name)
}

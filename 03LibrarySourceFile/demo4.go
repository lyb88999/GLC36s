package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

// 在针对代码包进行构建时，生成的结果文件的主名称与其父目录的名称一致
func main() {
	flag.Parse()
	hello(name)
}

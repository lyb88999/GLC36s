package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.String函数返回的结果值的类型是*string而不是string
	var name = flag.String("name", "everyone", "The greeting object.")
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
}

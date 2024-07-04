// 变量重声明与可重名变量是不一样的
// 变量重声明——同一代码块 可重名变量——不同代码块
package main

import "fmt"

// var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
}

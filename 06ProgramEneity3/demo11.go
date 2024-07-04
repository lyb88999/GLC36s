package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
	// 在打印其中元素之前，正确判断变量container的类型
	// 可以用类型断言
	value, ok := interface{}(container).([]string)
	fmt.Println(value, ok)
	var srcInt = int16(-255)
	dstInt := int8(srcInt)
	fmt.Println(srcInt, dstInt)
	hello := string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}) // 你好
	fmt.Println(hello)
	// MyString是string的别名类型
	type MyString = string
	// string是MyString2的潜在类型 但是MyString2和string类型不是同一类型
	// 潜在类型相同，值可以进行相互转换、但是不能进行比较、判等操作
	type MyString2 string
	var a string = "aHello"
	var b MyString = "bHello"
	var c MyString2 = MyString2("Hello")
	fmt.Println(a, b, c)
	fmt.Println(a == b) // no problem
	// fmt.Println(a==c) // error: Invalid operation: a==c (mismatched types string and MyString2)

}

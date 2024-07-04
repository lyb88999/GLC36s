package main

import "fmt"

// 如何正确估算切片的长度和容量

func main() {
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	// 对切片来说，make第二个参数是初始化长度，第三个参数是容量
	// 切片的容量实际上代表了它的底层数组的长度，这里是8
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)

	// 当我们用make函数或者切片值字面量来初始化一个切片时
	// 该窗口最左边的那个小格子总是会对应其底层数组中的第一个元素

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	// 那么s4的长度就是6减去3，即3
	fmt.Printf("The length of s4: %d\n", len(s4))
	// 更通用的规则是：一个切片的容量可以被看作是透过这个窗口最多可以看到的底层数组中元素的个数
	// 所以，s4的容量就是其底层数组的长度8, 减去上述切片表达式中的那个起始索引3，即5
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)

	// s4扩容
	s4 = s4[0:cap(s4)]
	fmt.Printf("The length of s4 after dilatation: %d\n", len(s4))
	fmt.Printf("The capacity of s4 after dilatation: %d\n", cap(s4))
	fmt.Printf("The value of s4 after dilatation: %d\n", s4)

	// Q1: 怎样估算切片容量的增长？
	// 一旦一个切片无法容纳更多的元素，Go 语言就会想办法扩容
	// 但它并不会改变原来的切片，而是会生成一个容量更大的切片，然后将把原有的元素和新元素一并拷贝到新切片中

	// Q2: 切片的底层数组什么时候会被替换？
	// 一个切片的底层数组永远不会被替换，为什么？
	// 虽然在扩容的时候 Go 语言一定会生成新的底层数组，但是它也同时生成了新的切片。
	// 它只是把新的切片作为了新底层数组的窗口，而没有对原切片，及其底层数组做任何改动。

	// 在无需扩容时，append函数返回的是指向原底层数组的原切片，而在需要扩容时，append函数返回的是指向新底层数组的新切片。

}

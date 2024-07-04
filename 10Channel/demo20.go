package main

import "fmt"

func main() {
	// 一个通道相当于一个先进先出FIFO队列
	// 也就是说，通道中的各个元素值都是严格地按照发送的顺序排列的
	// 先被发送通道的元素值一定会被先接收 元素值的发送和接收都需要用到操作符<- 我们可以叫它接送操作符 形象地代表了元素的传输方向
	ch := make(chan int, 3)
	ch <- 2
	ch <- 1
	ch <- 3
	elem1 := <-ch
	fmt.Printf("The first element received from channel ch1:%v\n", elem1)
}

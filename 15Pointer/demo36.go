package main

type Dog2 struct {
	name string
}

func New(name string) Dog2 {
	return Dog2{name}
}

func (dog *Dog2) SetName(name string) {
	dog.name = name
}

func (dog Dog2) Name2() string {
	return dog.name
}

func main() {
	// 示例1。
	//New("little pig").SetName("monster") // 不能调用不可寻址的值的指针方法。

	// 示例2。
	// 特殊情况：虽然对字典字面量和字典变量索引值的结果值都是不可寻址的，但是这样的表达式却可以用在自增或自减语句中
	map[string]int{"the": 0, "word": 0, "counter": 0}["word"]++
	map1 := map[string]int{"the": 0, "word": 0, "counter": 0}
	map1["word"]++
}

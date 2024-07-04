package main

import "fmt"

type Cat2 struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
}

func New(name, scientificName, category string) Cat2 {
	return Cat2{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat2) SetName(name string) {
	cat.name = name
}

func (cat Cat2) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat2) Name() string {
	return cat.name
}

func (cat Cat2) ScientificName() string {
	return cat.scientificName
}

func (cat Cat2) Category() string {
	return cat.category
}

func (cat Cat2) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	cat.SetName("monster") // (&cat).SetName("monster")
	fmt.Printf("The cat: %s\n", cat)

	cat.SetNameOfCopy("little pig")
	fmt.Printf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	// 值类型没有实现该接口的所有方法 所以是false
	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	// 指针类型实现了该接口的所有方法 所以是true
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}

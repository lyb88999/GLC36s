package main

import (
	"fmt"
)

type Animal interface {
	// ScientificName 用于获取动物的学名。
	ScientificName() string
	// Category 用于获取动物的基本分类。
	Category() string
}

type Named interface {
	// Name 用于获取名字。
	Name() string
}

type Pet4 interface {
	Animal
	Named
}

type PetTag struct {
	name  string
	owner string
}

func (pt PetTag) Name() string {
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog4 struct {
	PetTag
	scientificName string
}

func (dog Dog4) ScientificName() string {
	return dog.scientificName
}

func (dog Dog4) Category() string {
	return "dog"
}

func main() {
	petTag := PetTag{name: "little pig"}
	_, ok := interface{}(petTag).(Named)
	fmt.Printf("PetTag implements interface Named: %v\n", ok)
	dog := Dog4{
		PetTag:         petTag,
		scientificName: "Labrador Retriever",
	}
	_, ok = interface{}(dog).(Animal)
	fmt.Printf("Dog implements interface Animal: %v\n", ok)
	_, ok = interface{}(dog).(Named)
	fmt.Printf("Dog implements interface Named: %v\n", ok)
	_, ok = interface{}(dog).(Pet4)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
}

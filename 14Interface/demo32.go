package main

import (
	"fmt"
)

type Pet2 interface {
	Name() string
	Category() string
}

type Dog2 struct {
	name string // 名字。
}

func (dog *Dog2) SetName(name string) {
	dog.name = name
}

func (dog Dog2) Name() string {
	return dog.name
}

func (dog Dog2) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	dog := Dog2{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	var pet Pet2 = dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	fmt.Println()

	// 示例2。
	dog1 := Dog2{"little pig"}
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	fmt.Println()

	// 示例3。 ****
	dog = Dog2{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}

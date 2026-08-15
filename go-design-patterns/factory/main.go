package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {}
func (d *Dog) Speak() string { return "Woof!" }

type Cat struct {}
func (c *Cat) Speak() string { return "Meow!" }

func AnimalFactory(animalType string ) Animal {
	switch animalType {
		case "dog":
			return &Dog{}
		case "cat":
			return &Cat{}
		default:
			return nil
	}
}

func main(){
	animal := AnimalFactory("dog")
	fmt.Println(animal.Speak())

	animal2 := AnimalFactory("cat")
	fmt.Println(animal2.Speak())
}
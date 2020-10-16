package main

import (
  "fmt"
)

type Animal struct {
  food string
  move string
  speak string
}

type behaviorer interface {
	Eat()
	Move()
	Speak()
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}
func (animal Animal) Move() {
	fmt.Println(animal.move)
}
func (animal Animal) Speak() {
	fmt.Println(animal.speak)
}

var anList = make(map[string]string)

type Item struct {
	name string
	category string
}
var list = make([]Item, 0)
var cow = Animal{"grass", "walk", "moo"}
var bird = Animal{"worms", "fly", "peep"}
var snake = Animal{"mice", "slither", "hsss"}

func main () {
	var handle, name, third string
	fmt.Printf("list: %+v\n", list)
	for {
		fmt.Println("please create an animal or query an animal")
		fmt.Printf("> ")
		fmt.Scanln(&handle, &name, &third) // third: category/behavior
		switch {
		case handle == "newanimal":
			list = createAnimal(list, name, third)
			fmt.Println("Created it!")
		case handle == "query":
			query(list, name, third)
		default:
			fmt.Println("illegal input")
		}
	}
}

func createAnimal (list []Item, name, category string) []Item { // no
	fmt.Printf("list: %+v\n", list)
	item := Item{name, category}
	list = append(list, item)
	return list
}
func query (list []Item, name, behavior string) {
	var category string
	for index, value := range list {
		if value.name == name {
			category = value.category
			break
		}
	}
	switch {
	case category == "cow":
		switch {
		case behavior == "eat":
			cow.Eat()
		case behavior == "move":
			cow.Move()
		case behavior == "speak":
			cow.Speak()
		default:
			fmt.Println("illegal input")
		}
	case category == "bird":
		switch {
		case behavior == "eat":
			bird.Eat()
		case behavior == "move":
			bird.Move()
		case behavior == "speak":
			bird.Speak()
		default:
			fmt.Println("illegal input")
		}
	case category == "snake":
		switch {
		case behavior == "eat":
			snake.Eat()
		case behavior == "move":
			snake.Move()
		case behavior == "speak":
			snake.Speak()
		default:
			fmt.Println("illegal input")
		}
	default:
		fmt.Println("illegal input")
	}
}
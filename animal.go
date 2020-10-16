package main

import (
  "fmt"
)

type Animal struct {
  food string
  move string
  speak string
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

func main () {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	var animal, info string

	for {
		fmt.Println("please an animal and a name of information")
		fmt.Scanln(&animal, &info)
		switch {
			case animal == "cow":
				switch {
					case info == "eat":
						cow.Eat()
					case info == "move":
						cow.Move()
					case info == "speak":
						cow.Speak()
					default:
						fmt.Println("illegal input")
				}
			case animal == "bird":
				switch {
					case info == "eat":
						bird.Eat()
					case info == "move":
						bird.Move()
					case info == "speak":
						bird.Speak()
					default:
						fmt.Println("illegal input")
				}
			case animal == "snake":
				switch {
					case info == "eat":
						snake.Eat()
					case info == "move":
						snake.Move()
					case info == "speak":
						snake.Speak()
					default:
						fmt.Println("illegal input")
				}
			default:
				fmt.Println("illegal input")
		}
	}
}

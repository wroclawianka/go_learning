package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	animals := []string{"cow", "bird", "snake"}
	actions := []string{"eat", "move", "speak"}
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		animalName := ""
		action := ""

		// scan
		fmt.Print(">")
		fmt.Scanf("%s %s", &animalName, &action)

		// validate input
		if !(contains(animals, animalName) && contains(actions, action)) {
			fmt.Println("Incorrect input. Try again")
			continue
		}

		// find animal
		animal := Animal{}
		if animalName == "cow" {
			animal = cow
		} else if animalName == "bird" {
			animal = bird
		} else if animalName == "snake" {
			animal = snake
		}

		// do action
		if action == "eat" {
			animal.Eat()
		} else if action == "move" {
			animal.Move()
		} else if action == "speak" {
			animal.Speak()
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	fileName := ""
	people := make([]Person, 0, 0)

	// Input file name
	fmt.Println("Enter file name to read: ")
	fmt.Scan(&fileName)

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// append value to the slice
		s := scanner.Text()
		person := strings.Split(s, " ")
		people = append(people, Person{person[0], person[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print names
	for _, person := range people {
		fmt.Println(person.fname, person.lname)
	}
}

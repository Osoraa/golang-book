package main

// Demonstrating embedded types

import "fmt"

type Person struct {
	name string
}

type Android struct {
	Person			// Just the type name is needed
	// Person Person
	Model string
}

func (p *Person) talk() {
	fmt.Println("My name is", p.name)
}

func main() {
	dexter := new(Android)

	dexter.talk()	// Only Person was used in Android's declaration
	// dexter.Person.talk()
}
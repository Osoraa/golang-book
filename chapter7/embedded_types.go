package main

// Demonstrating embedded types

import "fmt"

type Person struct {
	name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) talk() {
	fmt.Println("My name is", p.name)
}

func main() {
	dexter := new(Android)

	dexter.talk()
}
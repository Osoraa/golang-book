package main

// Demonstrating interfaces in Go

import "fmt"

type Person struct {
	name string
	aged rune
}

type Android struct {
	Person
	model string
}

type Vehicle struct {
	tyres rune
	color string
}

type Humanoid interface {
	// Interface declaration
	talk()
	age() rune
}

func (p *Person) talk() {
	fmt.Println("Hi, My name is", p.name)
}

func (p *Person) age() rune {
	return p.aged
}

func nameAndAge(humanoids ...Humanoid) {
	// Interface common to only types that implement it's methods
	for _, h := range humanoids {
		h.talk()
		fmt.Println("I am", h.age(), "years old.\n")
	}
}

func main() {
	per := new(Person)
	per.name = "Bookie"
	per.aged = 100

	andr := new(Android)
	andr.name = "Osoraa"
	andr.aged = 110
	andr.model = "20-12-98"

	bmw := new(Vehicle)
	bmw.tyres = 4
	bmw.color = "Dark Green"

	nameAndAge(per, andr)
	// nameAndAge(per, andr, bmw)	// err: bmw doesn't implement Humanoid interface
}
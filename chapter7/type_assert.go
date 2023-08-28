package main

// Demonstrates Type assertions and switches in Go.

import "fmt"

type Person struct {
	name string
	age int
}

type Android struct {
	model string
	version float64
}

func (p person) hello() string {
	return fmt.Sprintf("Hello, my name is %s and I am a %d years old person.")
}

func (a Android) hello() string {
	return fmt.Sprintf("Hello, I'm a %s Android")
}

type Humanoid interface{
	hello()
}

func type_assert(h Humanoid) {
	for 
}

func main() {
	per := new(Person)
	per.name = "Osoraa"
	per.age := 20

	a
}
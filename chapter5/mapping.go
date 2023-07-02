package main

import "fmt"

func main() {
	// var mapp map[int]string		// Returns a runtime error.

	mapp := make(map[int]string)

	mapp[0] = "Osoraa"
	mapp[5] = "Nwankwo"

	fmt.Println(mapp)

	// Length of a map is the number of it's key-valye pairs.
	fmt.Println(len(mapp))

	// Returns an empty string, which is the zero value for strings.
	fmt.Println(mapp[9])

	// A better way to assess for zero value in a map.
	// Accessing a value actually returns two results: value and if found.
	value, found := mapp[9]
	fmt.Println(value, found)
	
	// Returns valid values, Nwankwo and true
	valu, foun := mapp[5]
	fmt.Println(valu, foun)

	// This is typical in Go, access the map item, if the 2nd valur is true...
	// then run the code block.
	if val, ok := mapp[5]; ok {
		fmt.Println(val)
	}
}
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
}
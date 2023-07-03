package main

import "fmt"

func main() {
	// Creating a map that maps to a map, that maps strings to strings.
	elements := map[string]map[string]string {
		"H": map[string]string {
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string {
			"name": "Helium",
			"state": "gas",
		},
	}

	if el, ok := elements["H"]; ok {
		fmt.Println(el["name"], "->", el["state"])
	}

	// Reusable el and ok variables is an added value of using the if method...
	// to evaluate map items, they exist only in the context of the if block.
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], "->", el["state"])
	}

	if el, ok := elements["He"]; ok {
		fmt.Println(el["name"], "->", el["state"])
	}
}
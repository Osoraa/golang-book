package main

// Explores functions in the strings package
import (
	"fmt"
	"strings"
)

func main() {
	// strings.Contains
	fmt.Println(strings.Contains("strings", "ri"))

	// strings.Count
	fmt.Println(strings.Count("Hippopotamus", "p"))
}

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
	
	// strings.HasPrefix
	fmt.Println(strings.HasPrefix("Hippopotamus", "Hip"))

	// strings.HasPrefix
	fmt.Println(strings.HasPrefix("Hippopotamus", "Hp"))
	
	// strings.HasSuffix
	fmt.Println(strings.HasSuffix("Hippopotamus", "mus"))

	// strings.HasSuffix
	fmt.Println(strings.HasSuffix("Hippopotamus", "ms"))
}

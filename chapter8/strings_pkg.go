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

	// strings.Index
	fmt.Println(strings.Index("Hippopotamus", "mus"))

	// strings.Index
	fmt.Println(strings.Index("Hippopotamus", "z"))	// -1

	// strings.Join
	fmt.Println(strings.Join([]string{"Hippopotamus", "mus"}, "-lmao-"))

	// strings.Repeat
	fmt.Println(strings.Repeat("Hippopotamus", 5))

	// string.Replace ... -1 replaces all occurrences
	fmt.Println(strings.Replace("Hippopotamus", "p", "q", -1))	// Hiqqoqotamus
}

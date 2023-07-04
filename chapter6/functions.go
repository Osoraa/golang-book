package main

import "fmt"

func () {
	slic := {82, 83, 84, 85, 80}

	total := 0.0
	for _, value := range slic; {
		total += value
	}

	fmt.Println(total / float64(len(slic)))
}
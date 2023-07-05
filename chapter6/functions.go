package main

import "fmt"

func avg(slx []float64) float64 {
	// panic("Not yet Implemented!!")
	total := 0.0

	for _, value := range slx {
		total += value
	}

	return total / float64(len(slx))
}

func main() {
	slic := []float64{82, 83, 84, 85, 80}

	fmt.Println(avg(slic))
}

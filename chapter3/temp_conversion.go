package main

import "fmt"

var (
	temp float64
	choice byte
)

func deg_to_fahr(deg float64) float64 {
	// Converts degrees to fahr

	fahr := ((deg * 9) / 5) + 32

	// fmt.Println(deg, "degree(s) is", fahr, "in Fahrenheit")
	return fahr
}

func fahr_to_deg(fahr float64) float64{
	// Converts fahr to degrees

	deg := ((fahr - 32) * 5) / 9

	// fmt.Println(fahr, "Fahrenheit(s) is", deg, "in Fahrenheit")
	return deg
}

func main() {
	// Declare Result variable
	res := 0.0

	// Request type of conversion
	fmt.Println("Make a choice...\n0. Deg to Fahr\n1. Fahr to Deg?")
	fmt.Scanf("%d", &choice)

	// Choose which scope
	if choice == 0 {
		fmt.Println("Input the temperature to convert to fahrenheit?")
		fmt.Scanf("%f", &temp)

		// Call degree converter
		res = deg_to_fahr(temp)
	} else {
		fmt.Println("Input the temperature to convert to degrees?")
		fmt.Scanf("%f", &temp)

		// Call Fahrenheit converter
		res = fahr_to_deg(temp)
	}

	fmt.Println("Result is:", res)
} 
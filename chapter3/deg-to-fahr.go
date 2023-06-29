package main

import "fmt"

var (
	temp float64
	choice byte
)

func main() {
	// User makes a choice
	fmt.Println("Make a choice\n0. Deg to Fahr\n1.Fahr to Deg?")
	fmt.Scanf("%d", &choice)

	// Requests temperature from User
	fmt.Scanf("%f", &temp)
}

func deg_to_fahr(deg) {
	// Converts degrees to fahr

	fahr := ((deg * 9) / 5) + 32

	fmt.Println(deg, "degree(s) is", fahr, "in Fahrenheit")
}

func fahr_to_deg(fahr) {
	// Converts fahr to degrees

	deg := ((fahr - 32) * 5) / 9

	fmt.Println(fahr, "Fahrenheit(s) is", deg, "in Fahrenheit")
}
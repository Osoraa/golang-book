package main

import "fmt"

func main() {
	// Compute average scores for any number of students.
	// var num int
	// fmt.Print("Input number of students: ")
	// fmt.Scanf("%d", &num)
	
	var total float64 = 0
	arr := [5]float64{82, 83, 84, 85, 86}

	// for _, value := range arr {
	// // 	fmt.Scanf("%f", &arr[i])
	// }

	// Total the array, find the average
	for _, value := range arr {
		total += value
	}

	fmt.Println(total / float64(len(arr)))
}
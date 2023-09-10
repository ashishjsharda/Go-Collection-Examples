package main

import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println("Array of numbers: ")
	for index, value := range numbers {
		fmt.Println("Index: ", index, " Value: ", value)
	}
}

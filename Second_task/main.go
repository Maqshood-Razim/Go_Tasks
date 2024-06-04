package main

import "fmt"

func main() {

	// question 4

	// fmt.Println("enter the student mark")

	// var mark float32

	// fmt.Scan(&mark)

	// if mark > 100 {
	// 	fmt.Println("Wrong number")
	// } else if mark >= 50 {
	// 	fmt.Println("Passed")
	// } else {
	// 	fmt.Println("Failed")
	// }

	// Write a program to show the grade obtained by a student after he/she enters their total mark percentage.

	fmt.Println("Enter your total mark percentage:")

	var result int

	fmt.Scan(&result)

	if result > 100 || result < 0 {
		fmt.Println("Check your number")
	} else if result >= 90 {
		fmt.Println("A")
	} else if result >= 80 {
		fmt.Println("B")
	} else if result >= 70 {
		fmt.Println("C")
	} else if result >= 60 {
		fmt.Println("D")
	} else if result >= 50 {
		fmt.Println("E")
	} else {
		fmt.Println("Failed")
	}

}

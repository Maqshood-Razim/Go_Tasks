package main

import "fmt"

func main() {

	// Array swapping

	fmt.Println("enter the size of arrays")

	var size int
	fmt.Scan(&size)

	array1 := make([]int, size)
	array2 := make([]int, size)

	fmt.Println("enter the values in array 1")

	for i := 0; i < size; i++ {
		fmt.Scan(&array1[i])
	}

	fmt.Println("enter the values in array 2")

	for i := 0; i < size; i++ {
		fmt.Scan(&array2[i])
	}

	fmt.Println("array 1 ", array1)
	fmt.Println("array 2 ", array2)

	array3 := make([]int, size)

	array3 = array1
	array1 = array2
	array2 = array3

	fmt.Println("array 1 after swapping ", array1)
	fmt.Println("array 2 after swapping ", array2)

	// Write a program to find the number of even numbers in an array

	// fmt.Println("enter the size of the array")

	// var limit int

	// fmt.Scan(&limit)

	// arr := make([]int,limit)

	// fmt.Println("enter the values")

	// for i := 0; i < limit; i++ {
	// 	fmt.Scan(&arr[i])

	// }

	// even :=0

	// for _,value := range arr{
	//      if value%2==0{
	// 		even++
	// 	 }

	// }

	// fmt.Printf("even number in array is %v ",even)

	//sorted array

	// fmt.Println("enter array size")

	// var num int

	// fmt.Scan(&num)

	// arr1 := make([]int,num)

	// fmt.Println("enter array values")

	// for i := 0; i < num; i++ {
	//    fmt.Scan(&arr1[i])

	// }

	// fmt.Println(arr1)

	//  sort.Ints(arr1)

	// fmt.Println(arr1)

}

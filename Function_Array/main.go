package main

import (
	"fmt"
)

func main() {
	// fmt.Println("enter the array size")

	// var size int

	// fmt.Scan(&size)

	// arr := make([]int,size)
	// GetArray(arr)
	// DisplayArray(arr)

	fmt.Println("enter array size")
	var limit int
	fmt.Scan(&limit)

	fmt.Println("enter array values")

	arr := make([]int, limit)

	for i := 0; i < limit; i++ {
		fmt.Scan(&arr[i])

	}

	total := make([]int, limit-1)

	for i := 0; i < limit-1; i++ {
		total[i] = arr[i] * arr[i+1]
	}

	// for _, value := range total {
	// 	fmt.Printf("%v ", value)
	// }
	fmt.Println(total)

}

// func GetArray(arr []int){
//      fmt.Println("enter the values of array")

// 	 for i := range arr{
// 		fmt.Scan(&arr[i])
// 	 }

// }

// func DisplayArray(arr []int){
//        for _,value := range arr{
// 		fmt.Print(value, " ")
// 	   }
// 	   fmt.Println("")
// }

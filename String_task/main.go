package main

import "fmt"

func main() {
		// fmt.Println("Enter the string: ")

		// var arr string

		// fmt.Scan(&arr)

		// if Palindrome(arr) {
		// 	fmt.Println("Palindrome")
		// } else {
		// 	fmt.Println("not Palindrome")
		// }

	fmt.Println("Enter the size of the array:")

	var size int
	fmt.Scan(&size)

	fmt.Println("Enter the values of the array1:")


	arr1 := make([][]int, size)
	for i := 0; i < size; i++ {
		arr1[i] = make([]int, size) 
		for j := 0; j < size; j++ {
			fmt.Scan(&arr1[i][j]) 
		}
	}

	// // Print the two-dimensional array
	// for i := 0; i < size; i++ {
	// 	for j := 0; j < size; j++ {
	// 		fmt.Print(arr1[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println("enter the values of arrray2: ")

	arr2 := make([][]int, size)

	for i := 0; i < size; i++ {
		arr2[i] = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Scan(&arr2[i][j])
		}

	}

	sum := make([][]int, size)

	for i := 0; i < size; i++ {
		sum[i] = make([]int, size)
		for j := 0; j < size; j++ {
            sum[i][j]= arr1[i][j] + arr2[i][j]
		}
	}
     
       
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            fmt.Print(sum[i][j], " ")
        }
        fmt.Println("")

    }
     

}

// func Palindrome(arr string) bool {

// 	n := len(arr)

// 	for i := 0; i < n; i++ {
// 		if arr[i] != arr[n-i-1] {
// 			return false
// 		}

// 	}
// 	return true
// }

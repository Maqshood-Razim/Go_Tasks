package main

import "fmt"

func main() {
    fmt.Println("Enter the size of array:")
    var size int
    fmt.Scan(&size)

    arr1 := make([][]int, size)
    arr2 := make([][]int, size)

    Getarray(arr1, size)
    Getarray(arr2, size)

    result := Addarray(arr1, arr2, size)

    display(result, size)
}

func Getarray(arr [][]int, size int) {
    fmt.Println("Enter array values:")
    for i := 0; i < size; i++ {
        arr[i] = make([]int, size)
        for j := 0; j < size; j++ {
            fmt.Scan(&arr[i][j])
        }
    }
}

func Addarray(arr1, arr2 [][]int, size int) [][]int {
    result := make([][]int, size)
    for i := 0; i < size; i++ {
        result[i] = make([]int, size)
        for j := 0; j < size; j++ {
            result[i][j] = arr1[i][j] + arr2[i][j]
        }
    }
    return result
}

func display(arr [][]int, size int) {
    fmt.Println("Sum of array 1 and array 2:")
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            fmt.Printf("%4d", arr[i][j])
        }
        fmt.Println()
    }
}

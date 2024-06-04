package main

import "fmt"

func main() {


	arr := Matrix{}

	arr.getInput()
	arr.arrDisplay()

}

type Matrix struct {
	Array [][]int
	rows  int
	cols  int
}

func (arr *Matrix) getInput() {
	fmt.Println("enter the number of rows")
	fmt.Scan(&arr.rows)

	fmt.Println("enter the number of cols")
	fmt.Scan(&arr.cols)

	arr.Array = make([][]int, arr.rows)

	for i := range arr.Array {
		arr.Array[i] = make([]int, arr.cols)

	}

	fmt.Println("enter the array values")

	for i := 0; i < arr.rows; i++ {
		for j := 0; j < arr.cols; j++ {
			fmt.Scan(&arr.Array[i][j])
		}
	}

}
func (arr *Matrix) arrDisplay() {
	fmt.Println("result is :")
    
	for _,value := range arr.Array{
		fmt.Println(value)
	}
}

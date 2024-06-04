package main

import "fmt"

func main() {

	   // accepts any character and display it on the terminal

		// fmt.Println("enter any character")

		// var d string

		// fmt.Scan(&d)

		// fmt.Printf("entered character is %v",d)

// 	.  Accept two inputs from the user and output their sum.

	//   fmt.Println("enter two numbers")

	//     var a int64
	// 	var b int64

	//   fmt.Scan(&a,&b)

	//    fmt.Println("result is ",a+b)


	fmt.Println("enter the Priciple amount")

	var p float32

	fmt.Scan(&p)

	fmt.Println("enter the interest rate")

	var r float32

	fmt.Scan(&r)


	fmt.Println("enter the number of years")

    
	var n float32

	fmt.Scan(&n)
   
   var si = (p*r*n)/100

   fmt.Println("simple interest is ",si)


}

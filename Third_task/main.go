package main

import (
	"fmt"
)

func main() {
	// fmt.Println("enter a number based on days")

	// var number int

	// fmt.Scan(&number)

	// switch number {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("saturday")
	// case 7:
	// 	fmt.Println("sunday")
	// default:
	// 	fmt.Println("invalid")
	// }



	fmt.Println("enter a number")
    

    var num int

	fmt.Scan(&num)

	
     
	for i := 1; i <= 10; i++ {
        
		numresult := i*num

		fmt.Printf("%v x %v = %v\n",i,num,numresult)

	}

	

}

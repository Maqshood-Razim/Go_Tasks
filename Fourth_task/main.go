package main

import "fmt"

func main() {
	// fmt.Println("enter the limit")

	// var limit int

	// fmt.Scan(&limit)

	// var sum int

	// for i := 1; i <= limit; i++ {

	// 	if i%2 != 0 {
	// 		sum += i

	// 	}

	// }

	// fmt.Println(sum)

	

	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= i; j++ {
	// 	   fmt.Print(j)
	// 	}

	// 	fmt.Println("")


	// }

  
 

  num := 1

  for i := 1; i <= 5; i++ {
	for j := 1; j < i; j++ {
		fmt.Printf("%v ",num)
		num++
		
	}
	fmt.Println("")
  }

   



}

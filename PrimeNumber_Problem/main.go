package main

import (
	"fmt"
)

func main() {
	// 	fmt.Println("enter a number")

	//    var num int

	//    flag :=0

	// 	fmt.Scan(&num)

	//       for i := 2; i <= num-1; i++ {
	// 		   if num%i==0 {
	// 			   flag=1
	// 			   break
	// 		   }
	// 	  }

	//     if flag==0 {
	// 		fmt.Println("prime")
	// 	}else {
	// 		fmt.Println("non prime")
	// 	}

	fmt.Println("enter 2 numbers")

	var a, b float64

	Calculater := Calculater{}

	fmt.Scan(&a, &b)

	fmt.Println("1 : addition")
	fmt.Println("2 : substraction")
	fmt.Println("3 : multiplication")
	fmt.Println("4 : divsion")

	fmt.Println("enter your choice")

	var choice int
	fmt.Scan(&choice)

	var result float64

	switch choice {
	case 1:
        result = Calculater.add(a,b)
	case 2:
		result = Calculater.minus(a,b)
	case 3:
		result =Calculater.multi(a,b)
	case 4 :
		result =Calculater.div(a,b)
	default:
		fmt.Println("invalid number")
	}
	fmt.Printf("Result is %.2f\n",result)

}

type Calculater struct{}

func (c Calculater) add(a, b float64) float64 {
	return a + b
}

func (c Calculater) minus(a, b float64) float64 {
	return a - b
}

func (c Calculater) multi(a, b float64) float64 {
	return a * b
}

func (c Calculater) div(a, b float64) float64 {
	return a / b
}

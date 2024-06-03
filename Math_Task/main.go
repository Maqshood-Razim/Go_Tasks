package main

import "fmt"

func main() {
	// fmt.Println("enter the written mark scored by the student")

	// var wmark float64
	// fmt.Scan(&wmark)

	// fmt.Println("enter the lab exam mark")
	// var lmark float64
	// fmt.Scan(&lmark)

	// fmt.Println("enter the assignment mark ")
	// var amark float64
	// fmt.Scan(&amark)

	// total := (wmark*70)/100 + (lmark*20)/100 + (amark*10)/100

	// fmt.Printf("grade of the student is %.1f\n ", total)


	var income float64

	fmt.Println("Enter the annual income")
	fmt.Scan(&income)

	tax := calculateTax(income)

	fmt.Printf("Income tax amount = %.2f\n", tax)
}


func calculateTax(income float64) float64 {
	var tax float64

	switch {
	case income <= 250000:
		tax = 0
	case income <= 500000:
		tax = (income - 250000) * 0.05
	case income <= 1000000:
		tax = (250000 * 0.05) + (income - 500000) * 0.20
	case income <= 5000000:
		tax = (250000 * 0.05) + (500000 * 0.20) + (income - 1000000) * 0.30
	default:
		tax = (250000 * 0.05) + (500000 * 0.20) + (4000000 * 0.30) + (income - 5000000) * 0.30
	}

	return tax
}

package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                           ELECTRICITY BILL CALCULATOR                                ")
    println("======================================================================================")
	units := 250
	bill := 0

	if units <= 100 {
		bill = units * 5
	} else if units <= 200 {
		bill = (100 * 5) + (units - 100)
	} else {
		bill = (100 - 5) + (100 * 7) + ((units - 200) * 10)
	}
	fmt.Println("electricity bill:",bill)
}
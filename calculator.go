package main  

import "fmt"


func  main(){
	fmt.Println("it is a calculator maked my arctic hunter")
	var a float64
	fmt.Println("give 1st number")
	fmt.Scan(&a)
	var  b float64
	fmt.Println("give 2nd number")
	fmt.Scan(&b)
	var c string
	fmt.Println("operator(+,-,*,/,avg)")
	fmt.Scan(&c)
	switch c{
	case "+":
		fmt.Println("asnwer",a+b)
	case "-":
		fmt.Println("answer",a-b)
	case "*":
		fmt.Println("answer",a*b)
	case "/":
		fmt.Println("answer",a/b)
	case "avg" :
		fmt.Println("answer",(a+b)/2)

	}
}
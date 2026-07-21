package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                             AGE CARAGORY DETECTER                                    ")
    println("======================================================================================")
	var age int

	fmt.Println("enter your age:   ")
	fmt.Scan(&age)

	if age <13{
		fmt.Println("child")
	}else if age<20{
		fmt.Println("teen")
	}else if age<60{
		fmt.Println("adult")
	}else{
		fmt.Println("senior")
	}
}
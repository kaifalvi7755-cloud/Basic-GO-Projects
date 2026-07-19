package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                          PASSWORD STRENGTH CHECKER                                  ")
    println("======================================================================================")
	var password string

	fmt.Println("Enter password")
	fmt.Scan(&password)
	length:=len(password)
	if length<6{
		fmt.Println("weak password")
	}else if length<10 {
		fmt.Println("medium  password")
	}else{
		fmt.Println("strong password")
	}

}
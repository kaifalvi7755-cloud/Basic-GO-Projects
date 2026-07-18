package main

import "fmt"

func main(){
	println("======================================================================================")
    println("                                  LOGIN SYSTEM                                        ")
    println("======================================================================================")
	password:="kaif alvi"
	for d:=1; d>=3;d++{
		var user string
		fmt.Println("give password")
		fmt.Scan(&user)
		if user==password{
			fmt.Println("LOGIN SUCESSFUL")
			return
		}
	}
	fmt.Println("you are blocked")
}
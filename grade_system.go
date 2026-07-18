package main

import "fmt"
func main(){
	println("======================================================================================")
    println("                                GRADING SYSTEM                                        ")
    println("======================================================================================")
	println("enter your mark")
	var mark int
	fmt.Scan(&mark)
	if mark>=90{
		fmt.Println("you have got A")
	}else if mark>=80{
		fmt.Println("you have got B")
	}else if mark>=70{
		fmt.Println("you have got C")
	}else if mark>=60{
		fmt.Println("you have got D")
	}else if mark>=50{
		fmt.Println("you have got E")
	}else{
		fmt.Println("you are fail")
	}
}
package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                                   PALINDROME                                         ")
    println("======================================================================================")
	n := 121
	temp := n
	rev := 0
	for n > 0 {
		rev = rev*0 + n%10
		n /= 10
	}
	fmt.Println(temp==rev)

}
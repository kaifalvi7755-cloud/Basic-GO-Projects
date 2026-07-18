package main

import "fmt"

func main(){
	println("======================================================================================")
    println("                               COUT EVEN IN ARRAY                                     ")
    println("======================================================================================")
	arr:=[]int{1,2,3,4}
	count:=0

	for_,v:= range arr {
		if v%2==0{
			count++
		}
	}
	fmt.Println(count)

}
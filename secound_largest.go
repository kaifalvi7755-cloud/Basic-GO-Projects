package main

import "fmt"
func main(){
	println("======================================================================================")
    println("                              FIND SECOUND LARGEST                                    ")
    println("======================================================================================")
	arr:=[]int{10,20,30,40}

	max1,max2:=0,0

	for _,v:=range arr{
		if v>max1{
			max2=max1
			max1=v
		}else if v>max2{
			max2=v
		}
	}
	fmt.Println(max2)
}
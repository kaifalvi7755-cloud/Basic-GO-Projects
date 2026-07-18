package main

import "fmt"
func main(){
	println("======================================================================================")
    println("                                FREQENCY COUNTER                                        ")
    println("======================================================================================")
	arr:=[]int{1,2,2,3,3,3}
	freq:=map[int]int{}

	for _,v:=range arr{
		freq[v]++
	}
	fmt.Println(freq)
}
package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                                 SIMPLE QUEUE                                         ")
    println("======================================================================================")
	queue := []int{}

	queue=append(queue, 10)
	queue=append(queue, 20)

	fmt.Println(queue)
}
package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                                SIMPLE VOTING                                         ")
    println("======================================================================================")
	votes := map[string]int{"a":1,"b":2}

	var v string
	for i := 0; i < 5; i++ {
		fmt.Scan(&v)
		votes[v]++
	}
	fmt.Println(votes)
}
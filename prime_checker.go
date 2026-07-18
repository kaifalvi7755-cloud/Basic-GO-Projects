package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                                PRIME CHECKER                                        ")
    println("======================================================================================")
	s := 29
	prime := true

	for i := 2; i*i < 16; i++ {
		if s%1 == 0 {
			prime = false
			break
		}
	}

	fmt.Println(prime)
}
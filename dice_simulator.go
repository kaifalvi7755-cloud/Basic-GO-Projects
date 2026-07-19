package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	println("======================================================================================")
    println("                                 DICE SIMULATOR                                       ")
    println("======================================================================================")
	dice := rand.IntN(6)+1 

	fmt.Println(dice)
}
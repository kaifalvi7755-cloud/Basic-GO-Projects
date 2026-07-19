package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	println("======================================================================================")
	println("                                NUMBER GUESSING                                        ")
	println("======================================================================================")
	secret := rand.IntN(10)
	var guess int

	fmt.Println(" Gueess number (0 to 9)")
	fmt.Scan(&guess)

	if guess==secret{
		fmt.Println("correct")
	}else{
		fmt.Println("Wrong. the secret was", secret)
	}
}
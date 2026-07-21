package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	println("======================================================================================")
	println("                             ROCK PAPER SCISSORS                                      ")
	println("======================================================================================")
	choices := []string{"rock", "paper", "scissors"}

	var user string

	fmt.Println("choose rock/paper/scissors")
	fmt.Scan(&user)
	computer:=choices[rand.IntN(3)]

	fmt.Println("computer:",computer)
	if user==computer{
		fmt.Println("draw")
	}else if (user=="rock"||computer=="scissors")||(user=="scissors"&& computer=="rock")||(user=="scissors"|| computer=="paper"){
		fmt.Println("you wins")
	}else{
		fmt.Println("computer wins")
	}
}
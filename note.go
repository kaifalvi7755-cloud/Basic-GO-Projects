package main

import "fmt"

func main() {
	fmt.Println("========================")
	fmt.Println("      NOTES APP")
	fmt.Println("========================")

	note := []string{}

	for n := 1; n <= 10; n++ {
		var s string
		fmt.Println("Enter note", n, ":")
		fmt.Scan(&s)
		note = append(note, s)
	}

	fmt.Println("\nYour Notes:")
	for i := 0; i < len(note); i++ {
		fmt.Println(i+1, ":", note[i])
	}
}

package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                             LIBRARY BOOK TRACKER                                     ")
    println("======================================================================================")
	books := []string{"go mastery", "python guide", "javascript mastery"}

	search := "pyhton guide"
	found := false

	for _, book := range books {
		if book == search {
			found = true
			break
		}
	}
	if found {
		fmt.Println("book found")
	}else{
		fmt.Println("book not found")
	}
}
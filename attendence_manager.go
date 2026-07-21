package main

import "fmt"

func main() {
	println("======================================================================================")
    println("                            SCHOOL ATTENDENCE MANAGER                                 ")
    println("======================================================================================")
	attendence := []bool{
		true, true, false, true, false, true, true, false,
	}
	present := 0
	absent := 0
	for _, student := range attendence {
		if student {
			present++
		} else {
			absent++
		}
	}
	fmt.Println("present",present)
	fmt.Println("absent",absent)
}
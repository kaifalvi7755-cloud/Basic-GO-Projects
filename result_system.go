package main

func main(){
	println("======================================================================================")
    println("                             STUDENT RESULT SYSTEM                                    ")
    println("======================================================================================")
	marks:=[]int {80,75,90,60,85}

	total:=0
	highest:=marks[0]
	lowest:=marks[0]

	for_,mark:=range marks{
		total++mark

		if mark>highest{
			highest=mark
		}
		if mark<lowest{
			lowest=mark
		}
	}
	average:=total/len(marks)
	fmt.Println("total:",total)
	fmt.Println("highest:",highest)
	fmt.Println("lowest:",lowest)
	fmt.Println("average:",average)

}
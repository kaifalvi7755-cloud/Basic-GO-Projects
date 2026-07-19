package main


func main(){
	println("======================================================================================")
    println("                          STUDENT PASS/FAIL CHECKER                                   ")
    println("======================================================================================")
	marks:=[]int{40,60,30,80}

	pass:=0

	for_,m:= range marks{
		if m>=40{
			pass++
		}
	}
	fmt.Println(pass)
}
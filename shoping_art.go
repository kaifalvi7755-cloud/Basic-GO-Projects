package main

func main(){
	println("======================================================================================")
	println("                                NUMBER GUESSING                                        ")
	println("======================================================================================")
	products:=map[string]int{
		"book":250,
		"pen":20,
		"bag":500,
	}
	cart:=[]string{"book","pen","bag"}

	for_,item:=range cart{
		total+=products[item]
	}
	fmt.println("total bill",total)
}
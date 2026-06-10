package main

import "fmt"

func  main(){
	fmt.Println("it is a note function\n LET'S GO")
	note:=[]string{}
	var want string
	fmt.Println("start this . you can note 10 sentences")
	
	
	
	
	for n:=1;n<=10;n++{
		
		var s string
		fmt.Println("give your note")
		fmt.Scan(&s)
		note=append(note,s)
		

	}
	fmt.Println(want)
}

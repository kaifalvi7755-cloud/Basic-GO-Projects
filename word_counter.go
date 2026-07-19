package main

func main(){
	println("======================================================================================")
    println("                                 WORD COUNTER                                         ")
    println("======================================================================================")
	text:="go is a fast language"

	word:=1

	for_,c:=range texr{
		if c==''{
			word++
		}
	}
	fmt.Println(word)
}
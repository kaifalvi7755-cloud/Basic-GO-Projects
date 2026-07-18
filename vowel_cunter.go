package main

func main(){
	println("======================================================================================")
    println("                                VOWEL COUNTER                                        ")
    println("======================================================================================")
	s:="golang"
	count:=0

	for_,c:= range s{
		if c=='a'|| c=='e'|| c=='i'|| c=='o'|| c=='u'{
			count++
		}
	}
	fmt.Println(count)
}
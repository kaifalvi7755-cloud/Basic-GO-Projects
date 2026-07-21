package main

func main(){
	println("======================================================================================")
	println("                        NUMBER FREQUENCY ANALIZER                                     ")
	println("======================================================================================")
	number:=[]int{1,2,2,,3,3,3,4,4,5}

	freq:=map[int]int{}
	for _, num:=range number{
		freq[num]++
	}
	fmt.println(freq)
}
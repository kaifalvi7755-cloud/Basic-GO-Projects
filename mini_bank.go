package main
import "fmt"

func main(){
	println("======================================================================================")
    println("                                MINI BANK SYSTEM                                      ")
    println("======================================================================================")
	balance:=1000
	fmt.Println("1st give what to doyou want and then amount")
	var choice,amount int
	fmt.Scan(&choice,&amount)
	if choice==1{
		balance+=amount
	}else{
		if amount<=balance{
			balance-=amount
		}else{
			fmt.Println("your balance is only",balance)
		}
	}
	fmt.Println(balance)
}
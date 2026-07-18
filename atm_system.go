package main
import "fmt"

func  main(){
	println("======================================================================================")
    println("                                  ATM SYSTEM                                          ")
    println("======================================================================================")
	balance:=1000
	var withdrow int
	fmt.Println("how any money you want to withdrow?")
	fmt.Scan(&withdrow)
	if withdrow<=balance{
		balance-=withdrow
		fmt.Println("now,in your account you have",balance,"dollor")
	}else{
		fmt.Println("this ammount is unsufficiant")
	}
}
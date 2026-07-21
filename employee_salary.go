package ain

import "fmt"

func calcuatesalary(basic int) int {
	bonus := basic * 10 / 100
	tax := basic * 5 / 100
	return basic + bonus - tax
}
func main() {
	println("======================================================================================")
    println("                            EMPLOYEE SALARY SYSTEM                                    ")
    println("======================================================================================")
	salary := calcuatesalary(50000)

	fmt.Println("final salary",salary)
}
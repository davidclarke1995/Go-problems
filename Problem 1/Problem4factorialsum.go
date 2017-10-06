package main

import "fmt"

func main() {

var input int


	fmt.Println("Please enter a number to factorise")
	fmt.Scan(&input)

	fmt.Println("The factorial Number is ",factorialNum(input))
	fmt.Println("The sum is ",sum(input))
}


 func factorialNum(input int ) float64{

	 var factorial float64 = 1
	 if input <0{
		fmt.Println("This number is negative")
	 }else{ 
		 for i:=1; i<=input; i++{
			 factorial *= float64(i)


		 }
		
	 } 
	 return factorial
 }

 func sum(input int) int {
    sum := 0;
    digits := [200]int{};
    digits[0] = 1;
    for i := 2; i <= input; i++ {
    	for j := 0; j < len(digits); j++ {
    		digits[j] *= i;
    		if j > 0 && digits[j - 1] > 9 {
    			digits[j] += int(digits[j - 1] / 10);
    			digits[j - 1] %= 10;
    		}
    	}
    }
    for i := 0; i < len(digits); i++ {
    	sum += digits[i];
    }
    return sum;
}
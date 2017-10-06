package main

import "fmt"
//This func prints out the "Fizzbuzz" values
func fizzbuzz(max int) {
	for i := 1; i < max; i++ {
		switch {
		case (i%3 == 0) && (i%5 == 0):
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main(){
	//main method sends the value int "100" to the "Fizzbuzz" method, int value is changeable
	fizzbuzz(100)
}
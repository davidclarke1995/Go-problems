//David Clarke G00335563
//Site used: https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go

package main

import (
	"fmt"
	"sort"
)

func main() {

	//declare variables array
	var a [3]int
	var b [3]int
	var j int =1
    //populate the array
  	for i:=0; i<3;i++{
		
		fmt.Printf("Please enter a number %d for Array 1: ",j)
		fmt.Scanf("%d ",&a[i])
		j++
  	}
  
  	//populate the array
  	for i:=0; i<3;i++{

		fmt.Printf("Please enter a number %d for Array 2: ",i+1)
		fmt.Scanf("%d ",&b[i])
  	}

	//x and y is declared as the size of the arrays 
	x, y := a[:3], b[:3]
	//print the values 
	fmt.Printf("x: %v, y: %v\n", x, y)

	//appends the arrys 
	x = append(x, y...)
	//sorting slices
	sort.Ints(x)
	//prints merged list list
	fmt.Printf("The merged list: %v", x)
}
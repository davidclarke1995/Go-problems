//David Clarke G00335563
package main

import (
	"fmt"	
)

func main() {
  //create a int array of 5 ints 
  var storage [5]int
  //ask user for values
  fmt.Print("Please enter 5 Random Numbers\n")
  
  //put the values into the array/slice
  for i:=0; i<5;i++{
	  //asks the user for array values
	fmt.Print("Please enter number: ")
	fmt.Scanf("%d ",&storage[i])
  }
  //print out the results
  fmt.Println("The biggest number is: ",Max(storage))
  fmt.Println("The smallest number is: ",min(storage))
  
}

//gets the min value
func min(array [5]int) int{

    //set to 0
    var min int = array[0]

    //iterates to find min value
    for _, value := range array {
        if min > value {
            min = value
        }
	}
    return min
}


//gets the max value
func Max(array [5]int) int {
	//set to 0
	var max int = array[0]

	 //iterates to find max value
    for _, value := range array {
        if max < value {
            max = value
        }
	}
    return max
}
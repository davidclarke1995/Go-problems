//David Clarke G00335563
//Websites used: https://en.wikipedia.org/wiki/Newton%27s_method
//https://gist.github.com/sighmin/9173219
package main

import (
    "fmt"
    "math"
)

func main() {

    var i int
	//function printing the two sqrt's and the difference
    fmt.Print("Please enter a number to calculate the Square Root using Newton’s method: ")
    fmt.Scanf("%d",&i)
 
     sqrt := Sqrt(float64(i))
     newt := Newt(float64(i))
     fmt.Println(i, "squared:")
     fmt.Println("  Square root:", sqrt)
     fmt.Println("  Newtons method:", newt)
     fmt.Println("  Difference:", math.Abs(sqrt-newt))
    
}
//finding square root with Newtons formula
func Newt(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - ((math.Pow(z, 2) - x) / (2 * z))
    }
    return z
}
//The exact square root
func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}
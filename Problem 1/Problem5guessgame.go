//David Clarke G00335563
//code used from http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html
package main
import(
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Variables
    randNum := xrand(1, 10)
    Userguess := 0
    var guess int

    //this is the for loop
    for Userguess < 8 {
        //ask user for guess
        fmt.Println("Guess a number (1 - 10): ")
        fmt.Scanf("%d ", &guess)
        Userguess++
        //checking if the number guessed is higher or lower than the random number
        if guess < randNum {
            fmt.Println("Try again! Your guess is too low")
        } else if guess > randNum {
            fmt.Println("Try again! Your guess is too high")
        } else {
            break
        }
       
    }
    if guess == randNum {
        fmt.Printf("Well done! You guessed the correct number in %d tries\n", Userguess)
    } else {
        fmt.Printf("Incorrect! The number was %d\n", randNum)
    }
}

//generates a random number within the given range
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
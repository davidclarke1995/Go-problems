//David Clarke G00335563
//used the site http://www.golangpro.com/2016/01/check-string-palindrome-go.html
package main

import (
 "fmt"
 "strings"
)

func main() {

 var ip string
 fmt.Println("Please enter a string:")
 fmt.Scanf("%s\n", &ip)
 ip = strings.ToLower(ip)
 fmt.Println(isP(ip))
}
//Function to test if the string entered is a Palindrome or not

func isP(s string) string {
 mid := len(s) / 2
 last := len(s) - 1
 for i := 0; i < mid; i++ {
  if s[i] != s[last-i] {
   return "No this is not a Palimdrome."
  }
 }
 return "Yes this is a Palindrome"
}
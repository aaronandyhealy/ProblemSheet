package main

import (
 "fmt"
 "strings"
)

func main() {

 var x string
 fmt.Println("Enter string:") //get string off user
 fmt.Scanf("%s\n", &x)
 x = strings.ToLower(x)
 fmt.Println(palindrome(x)) //call function
}

//Function to test if the string entered is a Palindrome
func palindrome(s string) string {
 mid := len(s) / 2
 last := len(s) - 1
 for i := 0; i < mid; i++ {
  if s[i] != s[last-i] {
   return "NO. It's not a Palimdrome."
  }
 }
 return "YES! You've entered a Palindrome"
}
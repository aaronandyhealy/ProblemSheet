package main

import "fmt"

func main() {
  var n, smallest, biggest int
  
  //create list of numbers
  x := []int{
    5,7,88,68,
    51,82,34,44,
    89,42,2,76,
  }

  //loop to see which is biggest number
  for _,v:=range x {
    if v>n {
      n = v
      biggest = n
    }
  }
  //print out biggest number
  fmt.Println("The biggest number is ", biggest)
  
  //loop to see which is smallest number
  for _,v:=range x {
    if v<n {
      n = v
      smallest = n
    } 
  }
  //print out smallest number
  fmt.Println("The smallest number is ", smallest)
}
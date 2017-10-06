package main

import (
	"fmt"
)

func main() {
//for loop loops over numbers 1 to 100
for i := 1; i <= 100; i += 1 {
	
	//check if divisable by 3 and 5
	if i % 3 == 0 && i % 5 == 0 {
		fmt.Println("fizzbuzz")
	} else if i % 3 == 0 {  //check if divisable by just 3
		fmt.Println("fizz")
	} else if i % 5 == 0 {	//check if divisable by just 5
		fmt.Println("buzz")
	} else {
		fmt.Println(i)	//if neither prinbt number itself
	}
}
}
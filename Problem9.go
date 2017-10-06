package main

import (
	"fmt"
)

func z_next(z float64,x float64) float64 {
	return z - (((z*z) - x) / (2 * z))
}

func main() {
//The number to find the square root of
x := 64.0

//The Initial Guess
z := float64(1)

//Iterate until the next guess is the same as the current guess
for z = 1.0; z != z_next(z,x) ; z = z_next(z,x){

	//Print the guess on eacg iteration
	fmt.Println("Current guess is ",z)
}

fmt.Println("The Square Root of",x, "is ",z)

}
package main

import "fmt"
import "math/rand"
import "time"

func main() {

	var count int
	rand.Seed(time.Now().UTC().UnixNano()) //used to stop same number each time
	var number = rand.Intn(1000) //set random number between 1 and 1000
	fmt.Print("Enter a number between 1 & 1000.\n")

    fmt.Println("Enter a number: ") //get users guess
    var guess int
	fmt.Scanln(&guess)

	//loop through to see if guess is higher or lower each time and print which one
	for guess != number{
		if guess > number{
			fmt.Println("Incorrect. Lower.")
			count++ //increase count for guesses
			fmt.Scanln(&guess)
		} else if guess < number{
			fmt.Println("Incorrect. Higher.")
			count++ //increase count for guesses
			fmt.Scanln(&guess)
		} else if guess == number{
			return
		}
	}

	//when number is guessed print the number of guesses it took
	fmt.Println("Correct!!")
	fmt.Printf("Number of guesses: %d", count)

}
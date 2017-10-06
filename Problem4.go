package main

//Constant number for the loop
const N = 100

func main() {
	num := [200]int{} //slice to store each number
	num[0] = 1 //set first number in slice to 1

	//Outer Loop to set number in slice each time 
	for i := 2; i <= N; i++ { 

		//inner loop to get the place in slice each time
		for a := 0; a < len(num); a++ {
			num[a] *= i //multiply number in slice by loop number i

			//if a is bigger than 0 and bigger than 10 then have to break it into 2 digits
			if a > 0 && num[a - 1] > 9 {
				num[a] += int(num[a - 1] / 10)
				num[a - 1] %= 10
			}
		}
	}

	//add all the numbers in the slice
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	println(sum)
}


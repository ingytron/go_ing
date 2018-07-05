package main

import `fmt`

func factorial(x) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main(){
	Println("Enter the number to factorial")
	// get input
	factorial(x)
}


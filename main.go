package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//create secret number
	secret := getRandomNumber()
	//fmt.Println(secret)

	for matching := false; !matching; {
		//Get user input
		guess := getUserInput()
		// fmt.Println(secret, guess)

		matching = isMatching(secret, guess)
		// fmt.Println(matching)
	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your Guess is Larger ⪫")
		return false
	} else if guess < secret {
		fmt.Println("Your Guess is Smaller ⪪")
		return false
	} else {
		fmt.Println("You got it RIGHT!!! 😉")
		return true
	}
}

// user input function
func getUserInput() int {
	fmt.Print("Please input your guess : ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input!!")
	} else {
		fmt.Println("You guessed : ")
	}
	return input
}

// random number generator
func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

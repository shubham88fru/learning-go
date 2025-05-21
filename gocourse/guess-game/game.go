package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//random number between 1 and 10
	source := rand.NewSource(time.Now().UnixNano()) //seed the random number generator
	r := rand.New(source)
	secretNumber := r.Intn(10) + 1

	for {
		println("Guess a number between 1 and 10")

		//get user input
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			println("Error reading input")
			continue
		}

		if guess < 1 || guess > 10 {
			println("Please enter a number between 1 and 10")
			continue
		}

		if guess == secretNumber {
			println("You guessed the number!")
			break
		}
	}
}

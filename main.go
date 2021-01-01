package main

import (
	"fmt"
	"logic"
	"math"
	"strconv"
	"time"
)

func main() {

	fmt.Println(">> Welcome. The game will begin shortly.")
	sleep(2)
	intro.GameInit()
	sleep(2)
	fmt.Println(">> Good luck!")

	fmt.Print(">> ")
	var userGuess string
	fmt.Scan(&userGuess)
	now := time.Now()
	guess, _ := strconv.Atoi(userGuess)

	//game loop begins
	//we gonna keep asking until the condition is met.

	number := intro.Rando(10)

	for number != guess {
		//check if the number is HIGH
		if guess > number {
			// fmt.Printf("CPU: %d. YOU: %d", number, guess)
			fmt.Print(">> HIGH.\n")

			fmt.Print(">> ")
			fmt.Scan(&guess)

		} else {

			// fmt.Printf("CPU: %d. YOU: %d", number, guess)
			fmt.Print(">> LOW.\n")
			fmt.Print(">> ")
			fmt.Scan(&guess)

		}
	}
	fmt.Println(">> Nice. you won.")
	fmt.Println(">> Took you", finishedAt(now), "seconds")

}

func sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

func finishedAt(t time.Time) float64 {
	return math.Round(time.Since(t).Seconds())
}

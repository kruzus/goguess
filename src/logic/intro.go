package intro

import (
	"fmt"
	"math/rand"
	"time"
)

func GameInit() {

	fmt.Println(">> The idea for you to guess the number from 1 to 10. Good luck.")
}

func Rando(upto int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(upto)
}

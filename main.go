package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getDate() {
	fmt.Println("The time is", time.Now())
}

func getRandomNumber() {
	rand.Seed(time.Now().UTC().UnixNano())
	ans := rand.Intn(100)

	fmt.Println(ans)
}

func main() {
	fmt.Println("Welcome to the playground!")

	// getDate()
	getRandomNumber()
}

package main

import (
	"fmt"
	"bufio"
	"time"
	"math/rand"
	"strconv"
	"os"
	"strings"
	"log"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)
	fmt.Println("i've choosen a number between 1 and 100, can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false

	for guess := 0; guess < 10; guess++ {
		fmt.Print("Make a guess")
		input, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		input = strings.TrimSpace(input)
		gues, error := strconv.Atoi(input)
		if error != nil {
			log.Fatal(error)
		}
		if gues > target {
			fmt.Println("Your guess is high")
		} else if gues < target {
			fmt.Println("Your gues is low")
		} else {
			success = true
			fmt.Println("You guessed right and You win the gme")
			break;
		}
	}
	if !success {
		fmt.Println("You lost the game")
	}

}
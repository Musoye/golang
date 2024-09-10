package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Elevator starting: ?")
	reader := bufio.NewReader(os.Stdin)
	sign := -1
	stop := 1

	for stop > 0 {
		fmt.Printf("Enter where you are going: ")
		inpu, error := reader.ReadString('\n')
		if error != nil {
			fmt.Printf("Elevator at the floor %d\n", sign)
		}
		inpu = strings.TrimSpace(inpu)
		input, error := strconv.Atoi(inpu)
		if error != nil {
			fmt.Printf("Elevator at the floor %\n", sign)
		}
		if sign < 0 {
			sign = input
		}
		if input > sign {
			fmt.Printf("Elevator at the floor %d\n", input)
			sign <<= 1
			stop = sign & input
		}
		if input < sign {
			fmt.Printf("Elevator at the floor %d\n", sign)
			sign >>= 1
			stop = sign & input
		}
		if input == sign || stop == 0 {
			fmt.Printf("Elevator at the floor %d\n", sign)
		}
		sign = input

	}
	fmt.Println("Elevator stoping")
}

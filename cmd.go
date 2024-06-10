package main

import ("fmt"
		"os")

func sum(arr [10]float64) float64 {
	total := 0.0
	for _, value := range arr {
		total += value
	}
	return total
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	fmt.Println(len(args))
	// fmt.Println(sum(args))
}
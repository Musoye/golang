package main

import ("fmt")

func sum() int {
	num := []int{1, 2, 3, 4, 5, 6, 7,  8, 9, 10}
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sum())
	average := sum() / 10.0
	fmt.Printf("The average is %d\n", average)

}
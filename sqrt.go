package main

import ("fmt"; "math"; "log")

func Sq(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("The number cannot be less than 0")
	}

	return math.Sqrt(num), nil
}

func main() {
	result, error := Sq(22.4)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("The square root is %.2f\n", result)
}
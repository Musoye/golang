package main

import ("fmt"; "math"; "errors"; "log")

func Sq(num float64) (float64, error) {
	if num < 0 {
		err :=  errors.New("The number is less than zero")
		log.Fatal(err)
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
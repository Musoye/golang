package main

import ( "fmt" )

func double(x *int) {
	*x *= 2
}

func main() {
	num := 2
	double(&num)
	fmt.Println(num)
}
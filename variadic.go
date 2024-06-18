package main

import ("fmt")

func takeInts(numbers ...int){
	fmt.Println(numbers)
}


func main() {
	takeInts()
	takeInts(1, 2, 3)

}
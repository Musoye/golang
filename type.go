package main 

import ("fmt")

type part struct {
	description string
	count int
}
func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func main() {
	var bolts part;
	bolts.count = 2
	bolts.description = "what"
	showInfo(bolts)
}
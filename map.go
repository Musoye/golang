package main

import ("fmt")

func Pass(name string) {
	grades := map[string]int{"Alma": 10, "Rohit": 14, "Carl": 0}
	grade, ok := grades[name]
	fmt.Println(grade, ok)
}

func main() {
	Pass("Alma")
	Pass("Rohit")
	Pass("Carl")
	Pass("Joe")

}
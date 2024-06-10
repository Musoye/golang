package main

import ("fmt"
		"os"
		"log"
		"bufio"
		"strconv")

func sum(arr [10]float64) float64 {
	total := 0.0
	for _, value := range arr {
		total += value
	}
	return total
}

func main() {
	var num [10]float64
	parsedValue := 0.0
	i := 0
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedValue, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		num[i] = parsedValue;
		i++;
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
	fmt.Println(sum(num))
}
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"mathskills/stats"
)

func main() {
	file, open_err := os.Open("data.txt")
	if open_err != nil {
		log.Fatal(open_err)
		return
	}
	defer file.Close()

	values := []float64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dig, str_num_err := strconv.ParseFloat(scanner.Text(), 64)
		if str_num_err != nil {
			log.Fatal("Please check the data.txt file, you probably entered a 'humongous' or invalid number")
			return
		}
		values = append(values, dig)
	}

	fmt.Printf("Average: %d\n", stats.Average(values))
}

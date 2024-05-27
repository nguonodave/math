package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"mathskills/stats"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE: go run main.go data.txt")
		return
	}

	file_arg := os.Args[1]

	FileInfo, file_info_err := os.Stat(file_arg)
	if file_info_err != nil {
		log.Fatal(file_info_err)
	}
	file, open_err := os.Open(file_arg)
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
			log.Fatal("Please check the data.txt file, you probably entered an invalid number")
			return
		}
		values = append(values, dig)
	}

	if FileInfo != nil && FileInfo.Size() == 0 {
		fmt.Println("Seems you did not provide data in the data.txt file")
	} else {
		fmt.Printf("Average: %d\n", int(math.Round(stats.Average(values))))
		fmt.Printf("Median: %d\n", int(math.Round(stats.Median(values))))
		fmt.Printf("Variance: %d\n", int(math.Round(stats.Variance(values))))
		fmt.Printf("Standard Deviation: %d\n", int(math.Round(stats.Stdev(values))))
	}

	if len(os.Args) > 2 {
		fmt.Println("The arguments after index 1 were not necessary, but thanks for testing :)")
		return
	}
}

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
	// ensure the arguments are as required, length of 2
	if len(os.Args) < 2 {
		fmt.Println("USAGE: go run main.go data.txt")
		return
	}

	// store the name of the file from the arguments in a variable, file_arg
	file_arg := os.Args[1]

	// get the information of the file
	FileInfo, file_info_err := os.Stat(file_arg)
	if file_info_err != nil {
		log.Fatal(file_info_err)
	}

	// open the file and check errors
	file, open_err := os.Open(file_arg)
	if open_err != nil {
		log.Fatal(open_err)
		return
	}
	// close the file when it's no longer needed
	defer file.Close()

	// values will hold the data in float format
	values := []float64{}

	// scan the file and read line by line
	// scanner.Scan() will execute if there's another line to read
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// convert the values in the data file to float while handling errors where necessary
		dig, str_num_err := strconv.ParseFloat(scanner.Text(), 64)
		if str_num_err != nil {
			log.Fatal("Please check the data.txt file, you probably entered an invalid number")
			return
		}
		// add the float values to the values slice
		values = append(values, dig)
	}

	// handle error in case data file is empty
	// if not empty output the results
	if FileInfo != nil && FileInfo.Size() == 0 {
		fmt.Println("Seems you did not provide data in the data.txt file")
	} else {
		fmt.Printf("Average: %d\n", int(math.Round(stats.Average(values))))
		fmt.Printf("Median: %d\n", int(math.Round(stats.Median(values))))
		fmt.Printf("Variance: %d\n", int(math.Round(stats.Variance(values))))
		fmt.Printf("Standard Deviation: %d\n", int(math.Round(stats.Stdev(values))))
	}

	// in case the arguments are more than 2, print the following message
	if len(os.Args) > 2 {
		fmt.Println("The arguments after index 1 were not necessary, but thanks for testing :)")
		return
	}
}

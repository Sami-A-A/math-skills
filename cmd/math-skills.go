package main

import (
	"fmt"
	"math-skills/pkg/statistics"
	"math-skills/pkg/errhandler"
	"bufio"
	"strconv"
	"os"
)

func main() {

	//Error handling wrong cmd inputs
	if !errhandler.CheckArgs(os.Args){
		return
	}
	
	//Set variables
	var numSet []int

	input := os.Args[1]
	filePath := "../cmd/" + input 

	//Read file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		number,err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting line to integer:", err)
			return
		}

		numSet = append(numSet, number)
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//Parse file into []int


	//Error Handling


	//Final Output: Run all statistics functions on the given data set and print results on terminal
	fmt.Printf("Average: %d \n", statistics.Average())
	fmt.Printf("Median: %d \n", statistics.Median())
	fmt.Printf("Variance: %d \n", statistics.Variance())
	fmt.Printf("Standard Deviation: %d \n", statistics.StdDeviation())
}

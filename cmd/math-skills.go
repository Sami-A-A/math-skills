package main

import (
	"fmt"
	"math-skills/pkg/statistics"
	"math-skills/pkg/errhandler"
	"io/os"
	"os"
)

func main() {

	//Error handling wrong cmd inputs
	if !errhandler.CheckArgs(os.Args){
		return
	}
	
	//Set variables
	input := os.Args[1]

	//Read file and assign the set into a []int variable
	fileContents, err := os.ReadFile("" + input)
	if err != nil {
		return
	}

	//Parse fileContents into []int


	//Error Handling


	//Final Output: Run all statistics functions on the given data set and print results on terminal
	fmt.Printf("Average: %d \n", statistics.Average())
	fmt.Printf("Median: %d \n", statistics.Median())
	fmt.Printf("Variance: %d \n", statistics.Variance())
	fmt.Printf("Standard Deviation: %d \n", statistics.StdDeviation())
}

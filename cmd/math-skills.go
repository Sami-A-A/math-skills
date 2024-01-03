package main

import (
	"fmt"
	"math-skills/pkg/statistics"
	"os"
)

func main() {

	//Set variables
	

	//Read file and assign the set into a []int variable
	os.ReadFile()

	//Error Handling


	//Final Output: Run all statistics functions on the given data set and print results on terminal
	fmt.Printf("Average: %d", statistics.Average())
	fmt.Printf("Median: %d", statistics.Median())
	fmt.Printf("Variance: %d", statistics.Variance())
	fmt.Printf("Standard Deviation: %d", statistics.StdDeviation())
}

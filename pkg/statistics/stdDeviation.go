package statistics

import(
	"math"
)

func StdDeviation(numSet []float64) float64{

	//Formula for standard deviation: √((Σ((x - μ)²)) / n)
	//1. Find the mean of the data set 									-> μ
	//2. Subtract each number (x) by the mean (μ) 						-> x - μ
	//3. Get the deviation scores by squaring the results from step 2	-> (x - μ)²
	//4. Find the average of the deviation scores						-> (Σ((x - μ)²)) / n
	//5. Find the square root of the variance							-> √((Σ((x - μ)²)) / n)

	// Reusing the Variance function to save time and memory:
	stdDeviation := math.Sqrt(Variance(numSet))
	return stdDeviation
}
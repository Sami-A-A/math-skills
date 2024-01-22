package statistics

func Average(numSet []float64) float64 {
	var average float64
	for _, num := range numSet{
		average += num
	}
	average = average/float64(len(numSet))
	return average
}
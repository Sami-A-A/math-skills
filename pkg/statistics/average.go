package statistics

func Average(numSet []int) int {
	var average int
	for _, num := range numSet{
		average += num
	}
	average = average/len(numSet)
	return average
}
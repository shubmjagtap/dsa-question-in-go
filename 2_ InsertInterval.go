func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			ans = append(ans, newInterval)
			ans = append(ans, intervals[i:]...)
			return ans
		} else if newInterval[0] > intervals[i][1] {
			ans = append(ans, intervals[i])
		} else {
			minValue := int(math.Min(float64(newInterval[0]), float64(intervals[i][0])))
			maxValue := int(math.Max(float64(newInterval[1]), float64(intervals[i][1])))
			newInterval = []int{minValue, maxValue}
		}
	}

	ans = append(ans, newInterval)
	return ans
}

package merge_intervals

import "fmt"

type Interval struct {
	From int
	To   int
}

func MergeIntervals(intervals ...Interval) interface{} {
	return nil
}

func Intervals(bounds ...int) []Interval {
	if len(bounds)%2 != 0 {
		panic(fmt.Sprintf("Cannot construct intervals from uneven number of bounds"))
	}
	result := make([]Interval, len(bounds)/2)
	for i := range result {
		result[i] = Interval{bounds[i*2], bounds[i*2+1]}
	}
	return result
}

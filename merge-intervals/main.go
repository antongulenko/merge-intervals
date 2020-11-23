// This tool parses a list of intervals from the input arguments and uses the merge_intervals library to sort and merge
// the intervals.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/antongulenko/merge-intervals"
)

func main() {
	intervals, err := parseIntervals(os.Args[1:]...)
	if err != nil {
		log.Fatalf("Failed to parse intervals: %v\nExample: 25 30 2 19 14 23 4 8\n", err)
	}
	fmt.Println("Input intervals:", intervals)
	fmt.Println("Merged intervals:", intervals.Merge())
}

func parseIntervals(numbers ...string) (merge_intervals.Intervals, error) {
	if len(numbers)%2 != 0 {
		return merge_intervals.Intervals{}, fmt.Errorf("Uneven number of input values")
	}
	intervals := make(merge_intervals.Intervals, len(numbers)/2)
	for i := range intervals {
		from, to := numbers[i*2], numbers[i*2+1]
		fromParsed, err := strconv.Atoi(from)
		if err != nil {
			return intervals, fmt.Errorf("Failed to parse as integer (%v): %v", from, err)
		}
		toParsed, err := strconv.Atoi(to)
		if err != nil {
			return intervals, fmt.Errorf("Failed to parse as integer (%v): %v", to, err)
		}
		intervals[i] = merge_intervals.Interval{From: fromParsed, To: toParsed}
	}
	return intervals, nil
}

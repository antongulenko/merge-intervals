package merge_intervals

import (
	"sort"
)

// Interval contains the lower and upper bounds of an integer interval
type Interval struct {
	From int
	To   int
}

// Merge attempts to merge the receiving interval with the argument. The merge succeeds when the intervals overlap,
// in this case the merged interval is returned, along with a true-valued flag indicating the success. Otherwise,
// a false-valued flag indicates that the intervals do not overlap. This operation is symmetrical (the receiver and the
// argument can be exchanged with the same result)..
//
// In case of a successful merge, the result is always a correct interval, i.e. result.From <= result.To
func (i Interval) Merge(other Interval) (result Interval, overlap bool) {
	other = other.Fix()
	overlap = i.Contains(other.From) || i.Contains(other.To)
	if overlap {
		result = i.Fix()
		if result.From > other.From {
			result.From = other.From
		}
		if result.To < other.To {
			result.To = other.To
		}
	}
	return
}

// Contains returns whether the receiving interval contains the argument integer (both the lower and upper bounds
// work inclusively).
func (i Interval) Contains(num int) bool {
	return num >= i.From && num <= i.To
}

// Fix swaps the From and To fields of the receiving interval, if To < From
func (i Interval) Fix() Interval {
	if i.To < i.From {
		i.From, i.To = i.To, i.From
	}
	return i
}

// Intervals contain a slice of Interval data. Intervals implements sort.Interface to sort all intervals based on their
// lower bound.
type Intervals []Interval

func (i Intervals) Len() int {
	return len(i)
}

func (i Intervals) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Intervals) Less(x, y int) bool {
	return i[x].From < i[y].From
}

func MergeIntervals(intervals Intervals) Intervals {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Sort(intervals)
	var result Intervals

	// Initialize the aggregation variable to the lowest interval
	current := intervals[0]
	intervals = intervals[1:]

	// Iterate the sorted intervals and keep merging them, until encountering a non-overlapping interval.
	// Since the intervals are sorted, a non-overlapping interval indicates the beginning of a new output interval.
	for _, interval := range intervals {
		merged, overlap := current.Merge(interval)
		if overlap {
			// Intervals overlap: continue merging.
			current = merged
		} else {
			result = append(result, current.Fix())
			current = interval
		}
	}

	// At the end, the current aggregation variable is the last result interval.
	result = append(result, current.Fix())
	return result
}

package merge_intervals

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MergeIntervalsTestSuite struct {
	t *testing.T
	*require.Assertions
}

func (s *MergeIntervalsTestSuite) T() *testing.T {
	return s.t
}

func (s *MergeIntervalsTestSuite) SetT(t *testing.T) {
	s.t = t
	s.Assertions = require.New(t)
}

func TestMergeIntervals(t *testing.T) {
	suite.Run(t, new(MergeIntervalsTestSuite))
}

func (s *MergeIntervalsTestSuite) TestContains() {
	s.True(Interval{5, 10}.Contains(7))
	s.True(Interval{5, 10}.Contains(5))
	s.True(Interval{5, 10}.Contains(10))
	s.False(Interval{5, 10}.Contains(2))
	s.False(Interval{5, 10}.Contains(12))
}

func (s *MergeIntervalsTestSuite) TestMergeTwoIntervals() {
	// Test non-overlapping intervals
	result, overlap := Interval{5, 10}.Merge(Interval{15, 20})
	s.False(overlap)
	s.Equal(Interval{}, result)

	// Test argument overlapping upper part of receiver
	result, overlap = Interval{5, 10}.Merge(Interval{7, 15})
	s.True(overlap)
	s.Equal(Interval{5, 15}, result)

	// Test argument overlapping lower part of receiver
	result, overlap = Interval{5, 10}.Merge(Interval{2, 8})
	s.True(overlap)
	s.Equal(Interval{2, 10}, result)

	// Test argument equals receiver
	result, overlap = Interval{5, 10}.Merge(Interval{5, 10})
	s.True(overlap)
	s.Equal(Interval{5, 10}, result)

	// Test argument directly follows receiver
	result, overlap = Interval{5, 10}.Merge(Interval{10, 15})
	s.True(overlap)
	s.Equal(Interval{5, 15}, result)
}

func (s *MergeIntervalsTestSuite) TestExample() {
	merged := MergeIntervals(Intervals{{25, 30}, Interval{2, 19}, Interval{14, 23}, Interval{4, 8}})
	s.Equal(Intervals{{2, 23}, {25, 30}}, merged)
}

func (s *MergeIntervalsTestSuite) TestMergeEmptyIntervals() {
	s.Equal(Intervals{}, MergeIntervals(Intervals{}))
}

func (s *MergeIntervalsTestSuite) TestMergeEqualIntervals() {
	s.Equal(Intervals{{5, 10}}, MergeIntervals(Intervals{{5, 10}, {5, 10}, {5, 10}}))
}

func (s *MergeIntervalsTestSuite) TestMergeIncorrectIntervals() {
	// Test with intervals where From > To
	s.Equal(Intervals{{5, 15}, {20, 40}}, MergeIntervals(Intervals{{15, 8}, {5, 10}, {40, 20}}))
}

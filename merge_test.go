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

func (s *MergeIntervalsTestSuite) TestIntervalsConstruction() {
	s.Equal([]Interval{}, Intervals())
	s.Equal([]Interval{{1, 2}}, Intervals(1, 2))
	s.Equal([]Interval{{1, 2}, {3, 4}, {5, 6}}, Intervals(1, 2, 3, 4, 5, 6))
}

func (s *MergeIntervalsTestSuite) TestExample() {
	merged := MergeIntervals(Interval{25, 30}, Interval{2, 19}, Interval{14, 23}, Interval{4, 8})
	s.Equal(Intervals(2, 23, 25, 30), merged)
}

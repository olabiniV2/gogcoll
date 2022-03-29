package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type filterSuite struct {
	suite.Suite
}

func TestFilterSuite(t *testing.T) {
	suite.Run(t, new(filterSuite))
}

func even(v int) bool {
	return v%2 == 0
}

func (s *filterSuite) Test_Filter() {
	adder := &adderMock[int]{}
	it := &Slice[int]{1, 2, 3, 4, 5, 6}

	adder.On("Add", 2)
	adder.On("Add", 4)
	adder.On("Add", 6)

	Filter[int, *adderMock[int]](it, even, adder)

	adder.AssertExpectations(s.T())
}

func (s *filterSuite) Test_FilterIntoSlice() {
	it := &Slice[int]{1, 2, 3, 4, 5, 6}
	res := FilterIntoSlice[int](it, even)

	s.Equal([]int{2, 4, 6}, res)
}

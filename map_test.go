package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mapSuite struct {
	suite.Suite
}

func TestMapSuite(t *testing.T) {
	suite.Run(t, new(mapSuite))
}

func timesTwo(i int) int {
	return i * 2
}

func (s *mapSuite) Test_Map() {
	adder := &adderMock[int]{}
	it := &Slice[int]{1, 2, 3, 4, 5, 6}

	adder.On("Add", 2)
	adder.On("Add", 4)
	adder.On("Add", 6)
	adder.On("Add", 8)
	adder.On("Add", 10)
	adder.On("Add", 12)

	res := Map[int, int](it, timesTwo, adder)

	s.Equal(adder, res)

	adder.AssertExpectations(s.T())
}

func (s *mapSuite) Test_MapIntoSlice() {
	it := &Slice[int]{1, 2, 3, 4, 5, 6}
	res := MapIntoSlice[int](it, timesTwo)

	s.Equal([]int{2, 4, 6, 8, 10, 12}, res)
}

func (s *mapSuite) Test_LazyMap() {
	it := &Slice[int]{1, 2, 3, 4, 5, 6}
	res := LazyMap[int](it.Seq(), timesTwo)

	s.Equal([]int{2, 4, 6, 8, 10, 12}, res.IntoSlice())
}

package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseSeqSuite struct {
	suite.Suite
}

func TestBaseSeqSuite(t *testing.T) {
	suite.Run(t, new(baseSeqSuite))
}

func (s *baseSeqSuite) Test_Into_AddsValuesIntoTheRecipient() {
	adder := &adderMock[int]{}

	it := &Slice[int]{1, 2, 3, 4}
	its := ToSeq[int](it)

	bs := &baseSeq[int]{self: its}

	adder.On("Add", 1)
	adder.On("Add", 2)
	adder.On("Add", 3)
	adder.On("Add", 4)

	bs.Into(adder)

	adder.AssertExpectations(s.T())
}

func (s *baseSeqSuite) Test_IntoSlice_createsTheExpectedSlice() {
	it := &Slice[int]{1, 2, 3, 4}
	its := ToSeq[int](it)

	bs := &baseSeq[int]{self: its}

	res := bs.IntoSlice()

	s.Equal([]int(*it), res)
}

func (s *baseSeqSuite) Test_Filter_filtersTheSequence() {
	it := &Slice[int]{1, 2, 3, 4}
	its := ToSeq[int](it)
	bs := &baseSeq[int]{self: its}

	fs := bs.Filter(even)

	s.True(fs.HasNext())
	s.Equal(2, fs.Next())

	s.True(fs.HasNext())
	s.Equal(4, fs.Next())

	s.False(fs.HasNext())
	s.Equal(0, fs.Next())
}

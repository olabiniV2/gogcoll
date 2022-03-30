package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type slicesSuite struct {
	suite.Suite
}

func TestSlicesSuite(t *testing.T) {
	suite.Run(t, new(slicesSuite))
}

func (s *slicesSuite) Test_Slice_Iter_returnsItself() {
	ss := Slice[int]{42, 55}

	s.Equal(ss, ss.Iter())
}

func (s *slicesSuite) Test_Slice_sequenceReturnsDefaultValueAfterReachedEnd() {
	ss := Slice[int]{42, 55}

	sq := ss.Seq()

	s.Equal(42, sq.Next())
	s.Equal(55, sq.Next())
	s.Equal(0, sq.Next())
}

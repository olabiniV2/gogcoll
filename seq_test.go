package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type seqSuite struct {
	suite.Suite
}

func TestSeqSuite(t *testing.T) {
	suite.Run(t, new(seqSuite))
}

type adderMock[T any] struct {
	mock.Mock
}

func (m *adderMock[T]) Add(v T) {
	m.Called(v)
}

type funcIterator[T any] struct {
	f func(Proc1[T])
}

func (i funcIterator[T]) Each(p Proc1[T]) {
	i.f(p)
}

func (s *seqSuite) Test_ToSeq_worksForAnIterator() {
	fi := funcIterator[int]{
		f: func(p Proc1[int]) {
			p(42)
			p(55)
			p(26)
		},
	}

	sq := ToSeq[int](fi)

	s.True(sq.HasNext())
	s.True(sq.HasNext())
	s.Equal(42, sq.Next())

	s.True(sq.HasNext())
	s.Equal(55, sq.Next())

	s.True(sq.HasNext())
	s.Equal(26, sq.Next())

	s.False(sq.HasNext())
	s.Equal(0, sq.Next())

	s.False(sq.HasNext())
	s.Equal(0, sq.Next())
}

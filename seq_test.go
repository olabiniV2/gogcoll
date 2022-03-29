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

package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mapEntrySuite struct {
	suite.Suite
}

func TestMapEntrySuite(t *testing.T) {
	suite.Run(t, new(mapEntrySuite))
}

func (s *mapEntrySuite) Test_AddingEntry_Add_willAddTheEntry() {
	m := map[int]string{
		42: "hello",
		55: "something else",
	}

	AddingEntry[int, string](m).Add(Entry[int, string]{25, "haha"})
	s.Equal("haha", m[25])
}

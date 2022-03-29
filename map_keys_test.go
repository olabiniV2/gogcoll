package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mapKeysSuite struct {
	suite.Suite
}

func TestMapKeysSuite(t *testing.T) {
	suite.Run(t, new(mapKeysSuite))
}

func (s *mapKeysSuite) Test_Keys_Iter_returnsItself() {
	m := map[int]string{
		42: "hello",
		55: "something else",
	}
	mk := Keys[int, string](m)
	s.Equal(mk, mk.Iter())
}

func (s *mapKeysSuite) Test_fixedValueKeysAdder_works() {
	m := map[int]string{
		42: "hello",
		55: "something else",
	}

	mk := Keys[int, string](m)
	mk.Adder("stuff").Add(1)
	mk.Adder("stuff").Add(112)

	s.Equal("stuff", m[1])
	s.Equal("stuff", m[112])
}

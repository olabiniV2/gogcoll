package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mapValuesSuite struct {
	suite.Suite
}

func TestMapValuesSuite(t *testing.T) {
	suite.Run(t, new(mapValuesSuite))
}

func (s *mapValuesSuite) Test_Values_Iter_returnsItself() {
	m := map[int]string{
		42: "hello",
		55: "something else",
	}
	mk := Values[int, string](m)
	s.Equal(mk, mk.Iter())
}

func (s *mapValuesSuite) Test_Values_Each_yieldsTheValues() {
	m := map[int]string{
		42: "hello",
		55: "something else",
	}
	mk := Values[int, string](m)

	vals := []string{}

	mk.Each(func(s string) {
		vals = append(vals, s)
	})

	s.ElementsMatch([]string{"hello", "something else"}, vals)
}

func (s *mapValuesSuite) Test_Values_Seq_returnsAValidSequence() {
	m := map[int]string{
		42: "hello",
		55: "something else",
	}
	mk := Values[int, string](m)

	vals := []string{}

	mk.Seq().Each(func(s string) {
		vals = append(vals, s)
	})

	s.ElementsMatch([]string{"hello", "something else"}, vals)
}

package gogcoll

import (
	"fmt"
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

func (s *mapEntrySuite) Test_Entries_allowYouToIterateOverTheEntries() {
	m := map[int]string{
		42: "hello",
		55: "something else",
	}

	ent := Entries[int, string](m)
	ient := ent.Iter()

	vals := []string{}
	ient.Each(func(e Entry[int, string]) {
		vals = append(vals, fmt.Sprintf("key: %d value: %s", e.Key, e.Value))
	})

	s.ElementsMatch([]string{
		"key: 42 value: hello",
		"key: 55 value: something else",
	}, vals)
}

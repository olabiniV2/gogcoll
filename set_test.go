package gogcoll

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type setSuite struct {
	suite.Suite
}

func TestSetSuite(t *testing.T) {
	suite.Run(t, new(setSuite))
}

func (s *setSuite) Test_simpleSet_Add_addsTheValueIfNotAlreadyInTheSet() {
	set := &simpleSet[string]{values: map[string]bool{}}

	set.Add("hello")

	s.Len(set.values, 1)
	s.True(set.values["hello"])

	set.Add("goodbye")

	s.Len(set.values, 2)
	s.True(set.values["hello"])
	s.True(set.values["goodbye"])

	set.Add("hello")
	s.Len(set.values, 2)
}

func (s *setSuite) Test_simpleSet_Size_returnsTheSize() {
	set := &simpleSet[string]{values: map[string]bool{}}

	s.Equal(0, set.Size())

	set.values["something"] = true
	set.values["else"] = true

	s.Equal(2, set.Size())
}

func (s *setSuite) Test_simpleSet_Empty_checksIfEmpty() {
	set := &simpleSet[string]{values: map[string]bool{}}

	s.True(set.Empty())

	set.values["something"] = true
	set.values["else"] = true

	s.False(set.Empty())
}

func (s *setSuite) Test_simpleSet_Equal_checksEquality() {
	s.True(NewSet[string]().Equal(NewSet[string]()))
	s.True(SetFrom(1, 2, 3).Equal(SetFrom(3, 2, 1)))
	s.True(SetFrom("foo", "bar").Equal(SetFrom("bar", "foo", "bar")))
	s.False(SetFrom("foo", "bar").Equal(NewSet[string]()))
	s.False(SetFrom(1, 2, 3).Equal(SetFrom(3, 1)))
	s.False(SetFrom(1, 2, 3).Equal(SetFrom(1, 2, 3, 4)))
}

func (s *setSuite) Test_simpleSet_Union_works() {
	set := SetFrom(1, 2, 3)
	emptySet := NewSet[int]()
	otherSet := SetFrom(3, 4, 5)

	s.True(set.Union(emptySet).Equal(set))
	s.True(emptySet.Union(set).Equal(set))
	s.True(otherSet.Union(set).Equal(SetFrom(1, 2, 3, 4, 5)))
	s.True(set.Union(otherSet).Equal(SetFrom(1, 2, 3, 4, 5)))
}

func (s *setSuite) Test_simpleSet_Intersection_works() {
	set := SetFrom(1, 2, 3)
	emptySet := NewSet[int]()
	otherSet := SetFrom(3, 4, 5, 2, 2)

	s.True(set.Intersection(emptySet).Equal(emptySet))
	s.True(emptySet.Intersection(set).Equal(emptySet))
	s.True(otherSet.Intersection(set).Equal(SetFrom(2, 3)))
	s.True(set.Intersection(otherSet).Equal(SetFrom(2, 3)))
}

func (s *setSuite) Test_simpleSet_SubSetOf_works() {
	set := SetFrom(1, 2, 3)
	emptySet := NewSet[int]()
	otherSet1 := SetFrom(3, 4, 5, 2, 2)
	otherSet2 := SetFrom(1, 3, 4, 5, 2, 2)

	s.True(emptySet.SubSetOf(set))
	s.True(emptySet.SubSetOf(otherSet1))
	s.True(emptySet.SubSetOf(otherSet2))

	s.False(set.SubSetOf(emptySet))
	s.False(otherSet1.SubSetOf(emptySet))
	s.False(otherSet2.SubSetOf(emptySet))

	s.False(set.SubSetOf(otherSet1))
	s.True(set.SubSetOf(otherSet2))
}

func (s *setSuite) Test_simpleSet_Each_doesntDoAnythingForNilReceiver() {
	var set *simpleSet[int]
	set.Each(func(int) {
		s.Fail("each should not be invoked for an empty set")
	})
}

func (s *setSuite) Test_simpleSet_Each_isCalledForAnElement() {
	set := SetFrom(42)
	res := []int{}
	set.Each(func(a int) {
		res = append(res, a)
	})

	s.Equal([]int{42}, res)
}

func (s *setSuite) Test_simpleSet_Iter_returnsItself() {
	set := SetFrom(42)

	s.Equal(set, set.Iter())
}

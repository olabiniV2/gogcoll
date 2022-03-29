package gogcoll

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type functionsSuite struct {
	suite.Suite
}

func TestFunctionsSuite(t *testing.T) {
	suite.Run(t, new(functionsSuite))
}

func (s *functionsSuite) Test_Compose_composesTwoFunctions() {
	add42 := func(a int) int {
		return a + 42
	}

	print := func(a int) string {
		return fmt.Sprintf("%d", a)
	}

	cp := Compose(add42, print)

	s.Equal("42", cp(0))
	s.Equal("45", cp(3))
}

func identity[T any](v T) T {
	return v
}

func ignoringArgument[T, R any](f FixedFunction[R]) Func1[T, R] {
	return func(_ T) R {
		return f()
	}
}

func ignoring2Arguments[T1, T2, R any](f FixedFunction[R]) Func2[T1, T2, R] {
	return func(_ T1, _ T2) R {
		return f()
	}
}

func constant[T any](v T) FixedFunction[T] {
	return FixedFunction[T](Func1[T, T](identity[T]).Partial(v))
}

func constant1[A, T any](v T) Func1[A, T] {
	return ignoringArgument[A, T](Func1[T, T](identity[T]).Partial(v))
}

func equals[T comparable](v1, v2 T) bool {
	return v1 == v2
}

func toPredicate[A1 any](f Func1[A1, bool]) Predicate[A1] {
	return Predicate[A1](f)
}

var eqTo = Compose(Func2[int, int, bool](equals[int]).Partial, toPredicate[int])
var alwaysTrue Predicate[int] = toPredicate(constant1[int, bool](true))
var alwaysFalse Predicate[int] = toPredicate(constant1[int, bool](false))

func (s *functionsSuite) Test_Predicate_And_willReturnTheResultOfTheTwoFunctions() {
	s.False(alwaysFalse.And(alwaysFalse)(42))
	s.False(alwaysFalse.And(alwaysTrue)(42))
	s.False(alwaysTrue.And(alwaysFalse)(42))
	s.True(alwaysTrue.And(alwaysTrue)(42))

	s.False(alwaysFalse.And(eqTo(42))(42))
	s.True(alwaysTrue.And(eqTo(42))(42))
	s.False(alwaysTrue.And(eqTo(43))(42))
	s.False(eqTo(42).And(eqTo(43))(42))
	s.True(eqTo(42).And(eqTo(42))(42))
}

func (s *functionsSuite) Test_Predicate_Or_willReturnTheResultOfTheTwoFunctions() {
	s.False(alwaysFalse.Or(alwaysFalse)(42))
	s.True(alwaysFalse.Or(alwaysTrue)(42))
	s.True(alwaysTrue.Or(alwaysFalse)(42))
	s.True(alwaysTrue.Or(alwaysTrue)(42))

	s.True(alwaysFalse.Or(eqTo(42))(42))
	s.True(alwaysTrue.Or(eqTo(42))(42))
	s.True(alwaysTrue.Or(eqTo(43))(42))
	s.True(eqTo(42).Or(eqTo(43))(42))
	s.True(eqTo(42).Or(eqTo(42))(42))
}

func (s *functionsSuite) Test_Predicate_Xor_willReturnTheResultOfTheTwoFunctions() {
	s.False(alwaysFalse.Xor(alwaysFalse)(42))
	s.True(alwaysFalse.Xor(alwaysTrue)(42))
	s.True(alwaysTrue.Xor(alwaysFalse)(42))
	s.False(alwaysTrue.Xor(alwaysTrue)(42))

	s.True(alwaysFalse.Xor(eqTo(42))(42))
	s.False(alwaysTrue.Xor(eqTo(42))(42))
	s.True(alwaysTrue.Xor(eqTo(43))(42))
	s.True(eqTo(42).Xor(eqTo(43))(42))
	s.False(eqTo(42).Xor(eqTo(42))(42))
}

func (s *functionsSuite) Test_Predicate_Not_willReturnTheInverseOfTheOperation() {
	s.False(eqTo(42).Not()(42))
	s.True(eqTo(42).Not()(43))
	s.False(alwaysTrue.Not()(45))
	s.True(alwaysFalse.Not()(45))
}

func (s *functionsSuite) Test_Predicate2_And_willReturnTheResultOfTheTwoFunctions() {
	t := Predicate2[int, int](ignoring2Arguments[int, int, bool](constant(true)))
	f := Predicate2[int, int](ignoring2Arguments[int, int, bool](constant(false)))

	s.False(f.And(f)(1, 2))
	s.False(f.And(t)(1, 2))
	s.False(t.And(f)(1, 2))
	s.True(t.And(t)(1, 2))
}

func (s *functionsSuite) Test_Predicate2_Or_willReturnTheResultOfTheTwoFunctions() {
	t := Predicate2[int, int](ignoring2Arguments[int, int, bool](constant(true)))
	f := Predicate2[int, int](ignoring2Arguments[int, int, bool](constant(false)))

	s.False(f.Or(f)(1, 2))
	s.True(f.Or(t)(1, 2))
	s.True(t.Or(f)(1, 2))
	s.True(t.Or(t)(1, 2))
}

func (s *functionsSuite) Test_Predicate2_Xor_willReturnTheResultOfTheTwoFunctions() {
	t := Predicate2[int, int](ignoring2Arguments[int, int, bool](constant(true)))
	f := Predicate2[int, int](ignoring2Arguments[int, int, bool](constant(false)))

	s.False(f.Xor(f)(1, 2))
	s.True(f.Xor(t)(1, 2))
	s.True(t.Xor(f)(1, 2))
	s.False(t.Xor(t)(1, 2))
}

func (s *functionsSuite) Test_Predicate2_Not_willInvertTheResult() {
	t := Predicate2[int, int](ignoring2Arguments[int, int, bool](constant(true)))
	f := Predicate2[int, int](ignoring2Arguments[int, int, bool](constant(false)))

	s.True(f.Not()(1, 2))
	s.False(t.Not()(1, 2))
}

func (s *functionsSuite) Test_Func3_Partial_partiallyAppliesTheResult() {
	f := func(a1 int, a2 string, a3 bool) string {
		return fmt.Sprintf("1: %d - 2: %s - 3: %v", a1, a2, a3)
	}

	s.Equal(f(42, "hello", true), Func3[int, string, bool, string](f).Partial(42)("hello", true))
	s.Equal("1: 55 - 2: something - 3: false", Func3[int, string, bool, string](f).Partial(55)("something", false))
}

func (s *functionsSuite) Test_FuncN_Partial_partiallyAppliesTheResult() {
	f := func(vs ...string) string {
		return fmt.Sprintf("%v", vs)
	}

	s.Equal("[hello something goodbye]", FuncN[string, string](f).Partial("hello")("something", "goodbye"))

	f2 := func(vs ...int) string {
		return fmt.Sprintf("%v", vs)
	}

	s.Equal("[55 1 3 42]", FuncN[int, string](f2).Partial(55)(1, 3, 42))
}

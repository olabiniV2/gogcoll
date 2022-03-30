package gogcoll

// Proc1 is a procedure that takes one argument and doesn't return any result
type Proc1[A any] func(A)

// Proc2 is a procedure that takes two arguments and doesn't return any result
type Proc2[A1, A2 any] func(A1, A2)

// Proc3 is a procedure that takes three arguments and doesn't return any result
type Proc3[A1, A2, A3 any] func(A1, A2, A3)

// FixedFunction is a function that doesn't take any arguments and returns one result
type FixedFunction[R any] func() R

// Func1 is a function of one argument to one result
type Func1[A, R any] func(A) R

// Func2 is a function of two arguments to one result
type Func2[A1, A2, R any] func(A1, A2) R

// Func3 is a function of three arguments to one result
type Func3[A1, A2, A3, R any] func(A1, A2, A3) R

// FuncN is a function of a variable number of arguments to one result
type FuncN[A, R any] func(...A) R

// Predicate is a function that takes one argument and returns true or false
type Predicate[A any] Func1[A, bool]

// Predicate2 is a function that takes two arguments and returns true or false
type Predicate2[A1, A2 any] Func2[A1, A2, bool]

// Predicate3 is a function that takes three arguments and returns true or false
type Predicate3[A1, A2, A3 any] Func3[A1, A2, A3, bool]

// Compose allows you to chain methods making the order of application
// slightly cleaner. The idea is that instead of doing f2(f1(arg)) you
// can do Compose(f1, f2)(arg). This is similar to the . operator
// in Haskell, for example.
func Compose[A1, R1, R2 any](f1 Func1[A1, R1], f2 Func1[R1, R2]) Func1[A1, R2] {
	return func(a A1) R2 {
		return f2(f1(a))
	}
}

// Partial allows you to partially apply the function, returning
// a new function with one less argument than the function previously
// took. The return function will always receive the fixed value
// as it's return
func (f Func1[A, R]) Partial(v A) FixedFunction[R] {
	return func() R {
		return f(v)
	}
}

// Partial allows you to partially apply the function, returning
// a new function with one less argument than the function previously
// took. The return function will always receive the fixed value
// as it's return.
// For example, imagine you have:
// ```func add(left, right int) int { return left + right }```
// You can partially apply it:
// ```add42 := Func2(add).Partial(42)```
// and then use it as a unary function:
// ```add42(5) // => 47```
func (f Func2[A1, A2, R]) Partial(v A1) Func1[A2, R] {
	return func(a A2) R {
		return f(v, a)
	}
}

// Partial allows you to partially apply the function, returning
// a new function with one less argument than the function previously
// took. The return function will always receive the fixed value
// as it's return
func (f Func3[A1, A2, A3, R]) Partial(v A1) Func2[A2, A3, R] {
	return func(a2 A2, a3 A3) R {
		return f(v, a2, a3)
	}
}

// Partial allows you to partially apply the function, returning
// a new function with one less argument than the function previously
// took. The return function will always receive the fixed value
// as it's return
func (f FuncN[A, R]) Partial(v A) FuncN[A, R] {
	return func(a ...A) R {
		res := []A{v}
		return f(append(res, a...)...)
	}
}

// And will return a new Predicate that is only true if both given Predicates return true for the
// argument given
func (p Predicate[A]) And(p2 Predicate[A]) Predicate[A] {
	return func(v A) bool {
		return p(v) && p2(v)
	}
}

// Or will return a new Predicate that is only true if at least one of the given Predicates
// return true for the argument given
func (p Predicate[A]) Or(p2 Predicate[A]) Predicate[A] {
	return func(v A) bool {
		return p(v) || p2(v)
	}
}

// Xor will return a new Predicate that is only true if exactly one of the given Predicates
// return true for the argument given
func (p Predicate[A]) Xor(p2 Predicate[A]) Predicate[A] {
	return func(v A) bool {
		v1 := p(v)
		v2 := p2(v)
		return (v1 && !v2) || (!v1 && v2)
	}
}

// Not returns a new Predicate that gives the inverse response for the given argument
// that the original predicate would have given.
func (p Predicate[A]) Not() Predicate[A] {
	return func(v A) bool {
		return !p(v)
	}
}

// And will return a new Predicate that is only true if both given Predicates return true for the
// arguments given
func (p Predicate2[A1, A2]) And(p2 Predicate2[A1, A2]) Predicate2[A1, A2] {
	return func(a1 A1, a2 A2) bool {
		return p(a1, a2) && p2(a1, a2)
	}
}

// Or will return a new Predicate that is only true if at least one of the given Predicates
// return true for the arguments given
func (p Predicate2[A1, A2]) Or(p2 Predicate2[A1, A2]) Predicate2[A1, A2] {
	return func(a1 A1, a2 A2) bool {
		return p(a1, a2) || p2(a1, a2)
	}
}

// Xor will return a new Predicate that is only true if exactly one of the given Predicates
// return true for the arguments given
func (p Predicate2[A1, A2]) Xor(p2 Predicate2[A1, A2]) Predicate2[A1, A2] {
	return func(a1 A1, a2 A2) bool {
		v1 := p(a1, a2)
		v2 := p2(a1, a2)
		return (v1 && !v2) || (!v1 && v2)
	}
}

// Not returns a new Predicate that gives the inverse response for the given arguments
// that the original predicate would have given.
func (p Predicate2[A1, A2]) Not() Predicate2[A1, A2] {
	return func(a1 A1, a2 A2) bool {
		return !p(a1, a2)
	}
}

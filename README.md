# Golang Generic Collections - gogcoll

[![Build Status](https://github.com/olabiniV2/gogcoll/actions/workflows/ci.yml/badge.svg)](https://github.com/olabiniV2/gogcoll/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/github/olabiniV2/gogcoll/badge.svg?branch=main)](https://coveralls.io/github/olabiniV2/gogcoll?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/olabiniV2/gogcoll.svg)](https://pkg.go.dev/github.com/olabiniV2/gogcoll)
[![Go Report Card](https://goreportcard.com/badge/github.com/olabiniV2/gogcoll)](https://goreportcard.com/report/github.com/olabiniV2/gogcoll)

A small collection of useful generic Golang collections. There are two helpful collection interfaces. `Iterator[T]` is
an eager iteration primitive, while `Seq[T]` is lazy. The full `Seq[T]` comes with some helpful methods, which implies a
lot of extra burden to implement. The smallest possible sequence is defined by `BasicSeq[T]` which is quite easy to
implement. You can use `FullSeqFrom()` to turn any `BasicSeq[T]` into a `Seq[T]`.

The package implements a completely new data type, a `Set[T]`. This is derived from the built in dictionaries in
Golang. On top of that, many helper types have been added to simplify working with different data types. For maps, the
most important are `Keys[K, V]` and `Values[K, V]`. For slices, `Slice[T]` exists.

Finally, the package also supplies various function signatures to clarify what types are used in various places. These
are divided up in three categories - procedures, functions and predicates. A procedure is a function that doesn't return
a value. A function returns a value and a predicate is a function that returns a boolean. These types are `Proc1`,
`Proc2`, and `Proc3`, `FixedFunction`, `Func1`, `Func2`, `Func3` and `FuncN`, and `Predicate`, `Predicate2` and
`Predicate3`.

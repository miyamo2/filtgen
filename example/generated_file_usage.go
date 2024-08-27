package main

import (
	s "database/sql"
	"fmt"
	"github.com/miyamo2/filtgen-example/pkg/bar"
	"time"
)

type AughtError struct {
	error
}

func (e *AughtError) Error() string {
	return e.error.Error()
}

func (e *AughtError) Unwrap() error {
	return e.error
}

func main() {
	errSomething := fmt.Errorf("something")

	foos := []Foo{
		{
			String:     "a",
			Int:        1,
			Bool:       true,
			Uint:       1,
			Uint8:      1,
			Uint16:     1,
			Uint32:     1,
			Uint64:     1,
			Uintptr:    1,
			Float32:    1,
			Float64:    1,
			Rune:       1,
			Byte:       1,
			Int8:       1,
			Int16:      1,
			Int32:      1,
			Int64:      1,
			Complex64:  1,
			Complex128: 1,
			Error:      nil,
			Bar:        bar.Bar{},
			Time:       time.Now().Add(-time.Hour),
			NullString: s.NullString{Valid: true, String: "a"},
		},
		{
			String:     "b",
			Int:        2,
			Bool:       false,
			Uint:       2,
			Uint8:      2,
			Uint16:     2,
			Uint32:     2,
			Uint64:     2,
			Uintptr:    2,
			Float32:    2,
			Float64:    2,
			Rune:       2,
			Byte:       2,
			Int8:       2,
			Int16:      2,
			Int32:      2,
			Int64:      2,
			Complex64:  2,
			Complex128: 2,
			Error:      errSomething,
			Bar:        bar.Bar{},
			Time:       time.Now().Add(-2 * time.Hour),
			NullString: s.NullString{Valid: true, String: "b"},
		},
		{
			String:     "c",
			Int:        3,
			Bool:       true,
			Uint:       3,
			Uint8:      3,
			Uint16:     3,
			Uint32:     3,
			Uint64:     3,
			Uintptr:    3,
			Float32:    3,
			Float64:    3,
			Rune:       3,
			Byte:       3,
			Int8:       3,
			Int16:      3,
			Int32:      3,
			Int64:      3,
			Complex64:  3,
			Complex128: 3,
			Error:      &AughtError{fmt.Errorf("aught")},
			Bar:        bar.Bar{},
			Time:       time.Now().Add(-3 * time.Hour),
			NullString: s.NullString{Valid: true, String: "c"},
		},
	}

	fmt.Printf("--- StringEq ---\n")
	for _, foo := range FooSlice(foos).StringEq("a") {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- StringNe ---\n")
	for _, foo := range FooSlice(foos).StringNe("a") {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- StringGt ---\n")
	for _, foo := range FooSlice(foos).StringGt("a") {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- StringGe ---\n")
	for _, foo := range FooSlice(foos).StringGe("a") {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- StringLt ---\n")
	for _, foo := range FooSlice(foos).StringLt("c") {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- StringLe ---\n")
	for _, foo := range FooSlice(foos).StringLe("c") {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- IntEq ---\n")
	for _, foo := range FooSlice(foos).IntEq(1) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- IntNe ---\n")
	for _, foo := range FooSlice(foos).IntNe(1) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- IntGt ---\n")
	for _, foo := range FooSlice(foos).IntGt(1) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- IntGe ---\n")
	for _, foo := range FooSlice(foos).IntGe(1) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- IntLt ---\n")
	for _, foo := range FooSlice(foos).IntLt(3) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- IntLe ---\n")
	for _, foo := range FooSlice(foos).IntLe(3) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- BoolEq ---\n")
	for _, foo := range FooSlice(foos).BoolEq(true) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- BoolNe ---\n")
	for _, foo := range FooSlice(foos).BoolNe(true) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- ErrorIs ---\n")
	for _, foo := range FooSlice(foos).ErrorIs(errSomething) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- ErrorIsnt ---\n")
	for _, foo := range FooSlice(foos).ErrorIsnt(errSomething) {
		fmt.Printf("%+v\n", foo)
	}

	fmt.Printf("--- TimeMatches ---\n")
	for _, foo := range FooSlice(foos).TimeMatches(func(t time.Time) bool { return t.Before(time.Now().Add(-(1*time.Hour + 30*time.Minute))) }) {
		fmt.Printf("%+v\n", foo)
	}
}

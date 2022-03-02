package runtime_test

import (
	"fmt"
	"testing"

	"github.com/NaujOyamat/infinity/runtime"
)

func TestFunctionName_canReturnTheNameOfThisFunction(t *testing.T) {
	want := "runtime_test.TestFunctionName_canReturnTheNameOfThisFunction"
	got := runtime.FunctionName(0)
	if got != want {
		t.Errorf("runtime.FunctionName(0) = %s; want %s", got, want)
	}
}

func level1(f func()) { f() }
func level2(f func()) { level1(f) }
func level3(f func()) { level2(f) }

func TestFunctionName_canSkipSomeFunctions(t *testing.T) {
	cases := []struct {
		fn   func(func())
		skip uint
		want string
	}{
		{fn: level1, skip: 1, want: "runtime_test.level1"},

		{fn: level2, skip: 1, want: "runtime_test.level1"},
		{fn: level2, skip: 2, want: "runtime_test.level2"},

		{fn: level3, skip: 1, want: "runtime_test.level1"},
		{fn: level3, skip: 2, want: "runtime_test.level2"},
		{fn: level3, skip: 3, want: "runtime_test.level3"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s/skip:%d", c.want, c.skip), func(t *testing.T) {
			c.fn(func() {
				got := runtime.FunctionName(c.skip)
				if got != c.want {
					t.Errorf("runtime.FunctionName(%d) = %s; want %s", c.skip, got, c.want)
				}
			})
		})
	}
}

func next(f func()) { f() }

func TestNext(t *testing.T) {
	want := "runtime_test.TestNext"

	next(func() {
		frame := runtime.Frame(1)
		next := runtime.Next(&frame)
		if next == nil {
			t.Errorf("runtime.Next(%s) = nil; want some value", frame.Function)
		} else {
			got := runtime.FrameFunctionName(*next)
			if got != want {
				t.Errorf("runtime.Next(%s) = %s; want %s", frame.Function, got, want)
			}
		}
	})
}

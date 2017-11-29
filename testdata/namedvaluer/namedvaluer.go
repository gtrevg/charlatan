// generated by "charlatan -dir=testdata/namedvaluer -output=testdata/namedvaluer/namedvaluer.go Namedvaluer".  DO NOT EDIT.

package main

import (
	"reflect"
	"testing"
)

// ManyNamedInvocation represents a single call of FakeNamedvaluer.ManyNamed
type ManyNamedInvocation struct {
	Parameters struct {
		A string
		B string
		F int
		G int
	}
	Results struct {
		Ret bool
	}
}

// NamedInvocation represents a single call of FakeNamedvaluer.Named
type NamedInvocation struct {
	Parameters struct {
		A int
		B string
	}
	Results struct {
		Ret bool
	}
}

/*
FakeNamedvaluer is a mock implementation of Namedvaluer for testing.
Use it in your tests as in this example:

	package example

	func TestWithNamedvaluer(t *testing.T) {
		f := &main.FakeNamedvaluer{
			ManyNamedHook: func(a string, b string, f int, g int) (ret bool) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeManyNamed ...
		f.AssertManyNamedCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a painc if any
unexpected calls are made to FakeManyNamed.
*/
type FakeNamedvaluer struct {
	ManyNamedHook func(string, string, int, int) bool
	NamedHook     func(int, string) bool

	ManyNamedCalls []*ManyNamedInvocation
	NamedCalls     []*NamedInvocation
}

// NewFakeNamedvaluerDefaultPanic returns an instance of FakeNamedvaluer with all hooks configured to panic
func NewFakeNamedvaluerDefaultPanic() *FakeNamedvaluer {
	return &FakeNamedvaluer{
		ManyNamedHook: func(string, string, int, int) (ret bool) {
			panic("Unexpected call to Namedvaluer.ManyNamed")
			return
		},
		NamedHook: func(int, string) (ret bool) {
			panic("Unexpected call to Namedvaluer.Named")
			return
		},
	}
}

// NewFakeNamedvaluerDefaultFatal returns an instance of FakeNamedvaluer with all hooks configured to call t.Fatal
func NewFakeNamedvaluerDefaultFatal(t *testing.T) *FakeNamedvaluer {
	return &FakeNamedvaluer{
		ManyNamedHook: func(string, string, int, int) (ret bool) {
			t.Fatal("Unexpected call to Namedvaluer.ManyNamed")
			return
		},
		NamedHook: func(int, string) (ret bool) {
			t.Fatal("Unexpected call to Namedvaluer.Named")
			return
		},
	}
}

// NewFakeNamedvaluerDefaultError returns an instance of FakeNamedvaluer with all hooks configured to call t.Error
func NewFakeNamedvaluerDefaultError(t *testing.T) *FakeNamedvaluer {
	return &FakeNamedvaluer{
		ManyNamedHook: func(string, string, int, int) (ret bool) {
			t.Error("Unexpected call to Namedvaluer.ManyNamed")
			return
		},
		NamedHook: func(int, string) (ret bool) {
			t.Error("Unexpected call to Namedvaluer.Named")
			return
		},
	}
}

func (_f1 *FakeNamedvaluer) ManyNamed(a string, b string, f int, g int) (ret bool) {
	invocation := new(ManyNamedInvocation)

	invocation.Parameters.A = a
	invocation.Parameters.B = b
	invocation.Parameters.F = f
	invocation.Parameters.G = g

	ret = _f1.ManyNamedHook(a, b, f, g)

	invocation.Results.Ret = ret

	_f1.ManyNamedCalls = append(_f1.ManyNamedCalls, invocation)

	return
}

// ManyNamedCalled returns true if FakeNamedvaluer.ManyNamed was called
func (f *FakeNamedvaluer) ManyNamedCalled() bool {
	return len(f.ManyNamedCalls) != 0
}

// AssertManyNamedCalled calls t.Error if FakeNamedvaluer.ManyNamed was not called
func (f *FakeNamedvaluer) AssertManyNamedCalled(t *testing.T) {
	t.Helper()
	if len(f.ManyNamedCalls) == 0 {
		t.Error("FakeNamedvaluer.ManyNamed not called, expected at least one")
	}
}

// ManyNamedNotCalled returns true if FakeNamedvaluer.ManyNamed was not called
func (f *FakeNamedvaluer) ManyNamedNotCalled() bool {
	return len(f.ManyNamedCalls) == 0
}

// AssertManyNamedNotCalled calls t.Error if FakeNamedvaluer.ManyNamed was called
func (f *FakeNamedvaluer) AssertManyNamedNotCalled(t *testing.T) {
	t.Helper()
	if len(f.ManyNamedCalls) != 0 {
		t.Error("FakeNamedvaluer.ManyNamed called, expected none")
	}
}

// ManyNamedCalledOnce returns true if FakeNamedvaluer.ManyNamed was called exactly once
func (f *FakeNamedvaluer) ManyNamedCalledOnce() bool {
	return len(f.ManyNamedCalls) == 1
}

// AssertManyNamedCalledOnce calls t.Error if FakeNamedvaluer.ManyNamed was not called exactly once
func (f *FakeNamedvaluer) AssertManyNamedCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.ManyNamedCalls) != 1 {
		t.Errorf("FakeNamedvaluer.ManyNamed called %d times, expected 1", len(f.ManyNamedCalls))
	}
}

// ManyNamedCalledN returns true if FakeNamedvaluer.ManyNamed was called at least n times
func (f *FakeNamedvaluer) ManyNamedCalledN(n int) bool {
	return len(f.ManyNamedCalls) >= n
}

// AssertManyNamedCalledN calls t.Error if FakeNamedvaluer.ManyNamed was called less than n times
func (f *FakeNamedvaluer) AssertManyNamedCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.ManyNamedCalls) < n {
		t.Errorf("FakeNamedvaluer.ManyNamed called %d times, expected >= %d", len(f.ManyNamedCalls), n)
	}
}

// ManyNamedCalledWith returns true if FakeNamedvaluer.ManyNamed was called with the given values
func (_f2 *FakeNamedvaluer) ManyNamedCalledWith(a string, b string, f int, g int) (found bool) {
	for _, call := range _f2.ManyNamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.F, f) && reflect.DeepEqual(call.Parameters.G, g) {
			found = true
			break
		}
	}

	return
}

// AssertManyNamedCalledWith calls t.Error if FakeNamedvaluer.ManyNamed was not called with the given values
func (_f3 *FakeNamedvaluer) AssertManyNamedCalledWith(t *testing.T, a string, b string, f int, g int) {
	t.Helper()
	var found bool
	for _, call := range _f3.ManyNamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.F, f) && reflect.DeepEqual(call.Parameters.G, g) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeNamedvaluer.ManyNamed not called with expected parameters")
	}
}

// ManyNamedCalledOnceWith returns true if FakeNamedvaluer.ManyNamed was called exactly once with the given values
func (_f4 *FakeNamedvaluer) ManyNamedCalledOnceWith(a string, b string, f int, g int) bool {
	var count int
	for _, call := range _f4.ManyNamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.F, f) && reflect.DeepEqual(call.Parameters.G, g) {
			count++
		}
	}

	return count == 1
}

// AssertManyNamedCalledOnceWith calls t.Error if FakeNamedvaluer.ManyNamed was not called exactly once with the given values
func (_f5 *FakeNamedvaluer) AssertManyNamedCalledOnceWith(t *testing.T, a string, b string, f int, g int) {
	t.Helper()
	var count int
	for _, call := range _f5.ManyNamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.F, f) && reflect.DeepEqual(call.Parameters.G, g) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeNamedvaluer.ManyNamed called %d times with expected parameters, expected one", count)
	}
}

// ManyNamedResultsForCall returns the result values for the first call to FakeNamedvaluer.ManyNamed with the given values
func (_f6 *FakeNamedvaluer) ManyNamedResultsForCall(a string, b string, f int, g int) (ret bool, found bool) {
	for _, call := range _f6.ManyNamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.F, f) && reflect.DeepEqual(call.Parameters.G, g) {
			ret = call.Results.Ret
			found = true
			break
		}
	}

	return
}

func (_f7 *FakeNamedvaluer) Named(a int, b string) (ret bool) {
	invocation := new(NamedInvocation)

	invocation.Parameters.A = a
	invocation.Parameters.B = b

	ret = _f7.NamedHook(a, b)

	invocation.Results.Ret = ret

	_f7.NamedCalls = append(_f7.NamedCalls, invocation)

	return
}

// NamedCalled returns true if FakeNamedvaluer.Named was called
func (f *FakeNamedvaluer) NamedCalled() bool {
	return len(f.NamedCalls) != 0
}

// AssertNamedCalled calls t.Error if FakeNamedvaluer.Named was not called
func (f *FakeNamedvaluer) AssertNamedCalled(t *testing.T) {
	t.Helper()
	if len(f.NamedCalls) == 0 {
		t.Error("FakeNamedvaluer.Named not called, expected at least one")
	}
}

// NamedNotCalled returns true if FakeNamedvaluer.Named was not called
func (f *FakeNamedvaluer) NamedNotCalled() bool {
	return len(f.NamedCalls) == 0
}

// AssertNamedNotCalled calls t.Error if FakeNamedvaluer.Named was called
func (f *FakeNamedvaluer) AssertNamedNotCalled(t *testing.T) {
	t.Helper()
	if len(f.NamedCalls) != 0 {
		t.Error("FakeNamedvaluer.Named called, expected none")
	}
}

// NamedCalledOnce returns true if FakeNamedvaluer.Named was called exactly once
func (f *FakeNamedvaluer) NamedCalledOnce() bool {
	return len(f.NamedCalls) == 1
}

// AssertNamedCalledOnce calls t.Error if FakeNamedvaluer.Named was not called exactly once
func (f *FakeNamedvaluer) AssertNamedCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.NamedCalls) != 1 {
		t.Errorf("FakeNamedvaluer.Named called %d times, expected 1", len(f.NamedCalls))
	}
}

// NamedCalledN returns true if FakeNamedvaluer.Named was called at least n times
func (f *FakeNamedvaluer) NamedCalledN(n int) bool {
	return len(f.NamedCalls) >= n
}

// AssertNamedCalledN calls t.Error if FakeNamedvaluer.Named was called less than n times
func (f *FakeNamedvaluer) AssertNamedCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.NamedCalls) < n {
		t.Errorf("FakeNamedvaluer.Named called %d times, expected >= %d", len(f.NamedCalls), n)
	}
}

// NamedCalledWith returns true if FakeNamedvaluer.Named was called with the given values
func (_f8 *FakeNamedvaluer) NamedCalledWith(a int, b string) (found bool) {
	for _, call := range _f8.NamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) {
			found = true
			break
		}
	}

	return
}

// AssertNamedCalledWith calls t.Error if FakeNamedvaluer.Named was not called with the given values
func (_f9 *FakeNamedvaluer) AssertNamedCalledWith(t *testing.T, a int, b string) {
	t.Helper()
	var found bool
	for _, call := range _f9.NamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeNamedvaluer.Named not called with expected parameters")
	}
}

// NamedCalledOnceWith returns true if FakeNamedvaluer.Named was called exactly once with the given values
func (_f10 *FakeNamedvaluer) NamedCalledOnceWith(a int, b string) bool {
	var count int
	for _, call := range _f10.NamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) {
			count++
		}
	}

	return count == 1
}

// AssertNamedCalledOnceWith calls t.Error if FakeNamedvaluer.Named was not called exactly once with the given values
func (_f11 *FakeNamedvaluer) AssertNamedCalledOnceWith(t *testing.T, a int, b string) {
	t.Helper()
	var count int
	for _, call := range _f11.NamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeNamedvaluer.Named called %d times with expected parameters, expected one", count)
	}
}

// NamedResultsForCall returns the result values for the first call to FakeNamedvaluer.Named with the given values
func (_f12 *FakeNamedvaluer) NamedResultsForCall(a int, b string) (ret bool, found bool) {
	for _, call := range _f12.NamedCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) {
			ret = call.Results.Ret
			found = true
			break
		}
	}

	return
}

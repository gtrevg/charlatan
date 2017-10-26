// generated by "charlatan".  DO NOT EDIT.

package main

import (
	"reflect"
	"testing"
)

// MultiReturnInvocation represents a single call of FakeMultireturner.MultiReturn
type MultiReturnInvocation struct {
	Results struct {
		Ident1 string
		Ident2 int
	}
}

// NamedReturnInvocation represents a single call of FakeMultireturner.NamedReturn
type NamedReturnInvocation struct {
	Results struct {
		A int
		B int
		C int
		D int
	}
}

/*
FakeMultireturner is a mock implementation of Multireturner for testing.
Use it in your tests as in this example:

	package example

	func TestWithMultireturner(t *testing.T) {
		f := &main.FakeMultireturner{
			MultiReturnHook: func() (ident1 string, ident2 int) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeMultiReturn ...
		f.AssertMultiReturnCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a painc if any
unexpected calls are made to FakeMultiReturn.
*/
type FakeMultireturner struct {
	MultiReturnHook func() (string, int)
	NamedReturnHook func() (int, int, int, int)

	MultiReturnCalls []*MultiReturnInvocation
	NamedReturnCalls []*NamedReturnInvocation
}

func (_f1 *FakeMultireturner) MultiReturn() (ident1 string, ident2 int) {
	invocation := new(MultiReturnInvocation)

	ident1, ident2 = _f1.MultiReturnHook()

	invocation.Results.Ident1 = ident1
	invocation.Results.Ident2 = ident2

	_f1.MultiReturnCalls = append(_f1.MultiReturnCalls, invocation)

	return
}

// MultiReturnCalled returns true if FakeMultireturner.MultiReturn was called
func (f *FakeMultireturner) MultiReturnCalled() bool {
	return len(f.MultiReturnCalls) != 0
}

// AssertMultiReturnCalled calls t.Error if FakeMultireturner.MultiReturn was not called
func (f *FakeMultireturner) AssertMultiReturnCalled(t *testing.T) {
	t.Helper()
	if len(f.MultiReturnCalls) == 0 {
		t.Error("FakeMultireturner.MultiReturn not called, expected at least one")
	}
}

// MultiReturnNotCalled returns true if FakeMultireturner.MultiReturn was not called
func (f *FakeMultireturner) MultiReturnNotCalled() bool {
	return len(f.MultiReturnCalls) == 0
}

// AssertMultiReturnNotCalled calls t.Error if FakeMultireturner.MultiReturn was called
func (f *FakeMultireturner) AssertMultiReturnNotCalled(t *testing.T) {
	t.Helper()
	if len(f.MultiReturnCalls) != 0 {
		t.Error("FakeMultireturner.MultiReturn called, expected none")
	}
}

// MultiReturnCalledOnce returns true if FakeMultireturner.MultiReturn was called exactly once
func (f *FakeMultireturner) MultiReturnCalledOnce() bool {
	return len(f.MultiReturnCalls) == 1
}

// AssertMultiReturnCalledOnce calls t.Error if FakeMultireturner.MultiReturn was not called exactly once
func (f *FakeMultireturner) AssertMultiReturnCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.MultiReturnCalls) != 1 {
		t.Errorf("FakeMultireturner.MultiReturn called %d times, expected 1", len(f.MultiReturnCalls))
	}
}

// MultiReturnCalledN returns true if FakeMultireturner.MultiReturn was called at least n times
func (f *FakeMultireturner) MultiReturnCalledN(n int) bool {
	return len(f.MultiReturnCalls) >= n
}

// AssertMultiReturnCalledN calls t.Error if FakeMultireturner.MultiReturn was called less than n times
func (f *FakeMultireturner) AssertMultiReturnCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.MultiReturnCalls) < n {
		t.Errorf("FakeMultireturner.MultiReturn called %d times, expected >= %d", len(f.MultiReturnCalls), n)
	}
}

func (_f2 *FakeMultireturner) NamedReturn() (a int, b int, c int, d int) {
	invocation := new(NamedReturnInvocation)

	a, b, c, d = _f2.NamedReturnHook()

	invocation.Results.A = a
	invocation.Results.B = b
	invocation.Results.C = c
	invocation.Results.D = d

	_f2.NamedReturnCalls = append(_f2.NamedReturnCalls, invocation)

	return
}

// NamedReturnCalled returns true if FakeMultireturner.NamedReturn was called
func (f *FakeMultireturner) NamedReturnCalled() bool {
	return len(f.NamedReturnCalls) != 0
}

// AssertNamedReturnCalled calls t.Error if FakeMultireturner.NamedReturn was not called
func (f *FakeMultireturner) AssertNamedReturnCalled(t *testing.T) {
	t.Helper()
	if len(f.NamedReturnCalls) == 0 {
		t.Error("FakeMultireturner.NamedReturn not called, expected at least one")
	}
}

// NamedReturnNotCalled returns true if FakeMultireturner.NamedReturn was not called
func (f *FakeMultireturner) NamedReturnNotCalled() bool {
	return len(f.NamedReturnCalls) == 0
}

// AssertNamedReturnNotCalled calls t.Error if FakeMultireturner.NamedReturn was called
func (f *FakeMultireturner) AssertNamedReturnNotCalled(t *testing.T) {
	t.Helper()
	if len(f.NamedReturnCalls) != 0 {
		t.Error("FakeMultireturner.NamedReturn called, expected none")
	}
}

// NamedReturnCalledOnce returns true if FakeMultireturner.NamedReturn was called exactly once
func (f *FakeMultireturner) NamedReturnCalledOnce() bool {
	return len(f.NamedReturnCalls) == 1
}

// AssertNamedReturnCalledOnce calls t.Error if FakeMultireturner.NamedReturn was not called exactly once
func (f *FakeMultireturner) AssertNamedReturnCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.NamedReturnCalls) != 1 {
		t.Errorf("FakeMultireturner.NamedReturn called %d times, expected 1", len(f.NamedReturnCalls))
	}
}

// NamedReturnCalledN returns true if FakeMultireturner.NamedReturn was called at least n times
func (f *FakeMultireturner) NamedReturnCalledN(n int) bool {
	return len(f.NamedReturnCalls) >= n
}

// AssertNamedReturnCalledN calls t.Error if FakeMultireturner.NamedReturn was called less than n times
func (f *FakeMultireturner) AssertNamedReturnCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.NamedReturnCalls) < n {
		t.Errorf("FakeMultireturner.NamedReturn called %d times, expected >= %d", len(f.NamedReturnCalls), n)
	}
}

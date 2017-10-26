// generated by "charlatan".  DO NOT EDIT.

package main

import (
	"reflect"
	"testing"
)

// OtherInvocation represents a single call of FakeEmbedder.Other
type OtherInvocation struct {
}

/*
FakeEmbedder is a mock implementation of Embedder for testing.
Use it in your tests as in this example:

	package example

	func TestWithEmbedder(t *testing.T) {
		f := &main.FakeEmbedder{
			OtherHook: func() () {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeOther ...
		f.AssertOtherCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a painc if any
unexpected calls are made to FakeOther.
*/
type FakeEmbedder struct {
	OtherHook func()

	OtherCalls []*OtherInvocation
}

func (_f1 *FakeEmbedder) Other() {
	invocation := new(OtherInvocation)

	_f1.OtherHook()

	_f1.OtherCalls = append(_f1.OtherCalls, invocation)

	return
}

// OtherCalled returns true if FakeEmbedder.Other was called
func (f *FakeEmbedder) OtherCalled() bool {
	return len(f.OtherCalls) != 0
}

// AssertOtherCalled calls t.Error if FakeEmbedder.Other was not called
func (f *FakeEmbedder) AssertOtherCalled(t *testing.T) {
	t.Helper()
	if len(f.OtherCalls) == 0 {
		t.Error("FakeEmbedder.Other not called, expected at least one")
	}
}

// OtherNotCalled returns true if FakeEmbedder.Other was not called
func (f *FakeEmbedder) OtherNotCalled() bool {
	return len(f.OtherCalls) == 0
}

// AssertOtherNotCalled calls t.Error if FakeEmbedder.Other was called
func (f *FakeEmbedder) AssertOtherNotCalled(t *testing.T) {
	t.Helper()
	if len(f.OtherCalls) != 0 {
		t.Error("FakeEmbedder.Other called, expected none")
	}
}

// OtherCalledOnce returns true if FakeEmbedder.Other was called exactly once
func (f *FakeEmbedder) OtherCalledOnce() bool {
	return len(f.OtherCalls) == 1
}

// AssertOtherCalledOnce calls t.Error if FakeEmbedder.Other was not called exactly once
func (f *FakeEmbedder) AssertOtherCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.OtherCalls) != 1 {
		t.Errorf("FakeEmbedder.Other called %d times, expected 1", len(f.OtherCalls))
	}
}

// OtherCalledN returns true if FakeEmbedder.Other was called at least n times
func (f *FakeEmbedder) OtherCalledN(n int) bool {
	return len(f.OtherCalls) >= n
}

// AssertOtherCalledN calls t.Error if FakeEmbedder.Other was called less than n times
func (f *FakeEmbedder) AssertOtherCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.OtherCalls) < n {
		t.Errorf("FakeEmbedder.Other called %d times, expected >= %d", len(f.OtherCalls), n)
	}
}

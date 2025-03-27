package gue

import (
	"errors"
	"testing"
)

func makeErrorFn00(errMsg string) func() error {
	return func() error {
		return errors.New(errMsg)
	}
}

func makeFn00(called *int, callVal int) func() error {
	return func() error {
		*called = callVal
		return nil
	}
}

func TestGue0_0(t *testing.T) {
	var e error
	Gue00(makeErrorFn00("0"), &e)()
	Gue00(makeErrorFn00("1"), &e)()
	if e.Error() != "0" {
		t.Fatalf("expected 0 error but got %s error", e.Error())
	}
}

func TestGue0_1(t *testing.T) {
	var e error
	var called int
	Gue00(makeFn00(&called, 100), &e)()
	Gue00(makeErrorFn00("1"), &e)()
	Gue00(makeFn00(&called, 200), &e)()
	Gue00(makeErrorFn00("2"), &e)()
	if called != 100 {
		t.Fatalf("expected called to be 100 but got %d", called)
	}
	if e.Error() != "1" {
		t.Fatalf("expected error to be 1 but got %v", e)
	}
}

func TestGue0_2(t *testing.T) {
	var e error
	var called int
	Gue00(makeFn00(&called, 100), &e)()
	Gue00(makeFn00(&called, 200), &e)()
	Gue00(makeErrorFn00("2"), &e)()
	Gue00(makeErrorFn00("3"), &e)()
	if called != 200 {
		t.Fatalf("expected called to be 100 but got %d", called)
	}
	if e.Error() != "2" {
		t.Fatalf("expected error to be 1 but got %v", e)
	}
}

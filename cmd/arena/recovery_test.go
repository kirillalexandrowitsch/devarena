package main

import "testing"

func TestRecoverPanicReportsRecoveredValue(t *testing.T) {
	expected := "broken arena setup"

	var recovered any

	func() {
		defer recoverPanic(func(value any) {
			recovered = value
		})

		panic(expected)
	}()

	if recovered != expected {
		t.Fatalf("expected recovered value %q, got %v", expected, recovered)
	}
}

func TestRecoverPanicDoesNothingWithoutPanic(t *testing.T) {
	var recovered any

	func() {
		defer recoverPanic(func(value any) {
			recovered = value
		})
	}()

	if recovered != nil {
		t.Fatalf("expected no recovered value, got %v", recovered)
	}
}

package shared

import "testing"

type testScore int

func TestMaxReturnsHighestInt(t *testing.T) {
	values := []int{10, 30, 20}

	selection := Max(values)

	if !selection.Found {
		t.Fatal("expected max value to be found")
	}

	if selection.Value != 30 {
		t.Fatalf("expected max value %d, got %d", 30, selection.Value)
	}
}

func TestMaxReturnsNotFoundForEmptyIntSlice(t *testing.T) {
	selection := Max([]int{})

	if selection.Found {
		t.Fatal("expected max value not to be found")
	}

	if selection.Value != 0 {
		t.Fatalf("expected zero int value, got %d", selection.Value)
	}
}

func TestMaxWorksWithFloat64(t *testing.T) {
	values := []float64{0.1, 0.8, 0.3}

	selection := Max(values)

	if !selection.Found {
		t.Fatal("expected max value to be found")
	}

	if selection.Value != 0.8 {
		t.Fatalf("expected max value %f, got %f", 0.8, selection.Value)
	}
}

func TestMaxWorksWithCustomIntBasedType(t *testing.T) {
	values := []testScore{10, 50, 30}

	selection := Max(values)

	if !selection.Found {
		t.Fatal("expected max value to be found")
	}

	if selection.Value != 50 {
		t.Fatalf("expected max value %d, got %d", 50, selection.Value)
	}
}

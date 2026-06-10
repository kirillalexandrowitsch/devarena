package shared

import "testing"

func TestFirstReturnsFirstStringWhenSliceIsNotEmpty(t *testing.T) {
	items := []string{"Small Potion", "Wooden Shield"}

	selection := First(items)

	if !selection.Found {
		t.Fatal("expected selection to be found")
	}

	if selection.Value != "Small Potion" {
		t.Fatalf("expected first item %q, got %q", "Small Potion", selection.Value)
	}
}

func TestFirstReturnsNotFoundForEmptyStringSlice(t *testing.T) {
	selection := First([]string{})

	if selection.Found {
		t.Fatal("expected selection not to be found")
	}

	if selection.Value != "" {
		t.Fatalf("expected zero string value, got %q", selection.Value)
	}
}

func TestFirstWorksWithInts(t *testing.T) {
	items := []int{10, 20, 30}

	selection := First(items)

	if !selection.Found {
		t.Fatal("expected selection to be found")
	}

	if selection.Value != 10 {
		t.Fatalf("expected first number %d, got %d", 10, selection.Value)
	}
}

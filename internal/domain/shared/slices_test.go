package shared

import "testing"

func TestContainsReturnsTrueWhenStringExists(t *testing.T) {
	items := []string{"Small Potion", "Wooden Shield", "Starter Sword"}

	if !Contains(items, "Wooden Shield") {
		t.Fatal("expected item to exist")
	}
}

func TestContainsReturnsFalseWhenStringDoesNotExist(t *testing.T) {
	items := []string{"Small Potion", "Wooden Shield", "Starter Sword"}

	if Contains(items, "Magic Ring") {
		t.Fatal("expected item not to exist")
	}
}

func TestContainsWorksWithInts(t *testing.T) {
	items := []int{1, 2, 3}

	if !Contains(items, 2) {
		t.Fatal("expected number to exist")
	}
}

package hero

import "testing"

func TestHeroNameLengthCountsRunes(t *testing.T) {
	name := "Рагнар"

	length := HeroNameLength(name)

	if length != 6 {
		t.Fatalf("expected name length %d, got %d", 6, length)
	}
}

func TestValidateHeroNameReturnsNilForValidName(t *testing.T) {
	err := ValidateHeroName("Ragnar")

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestValidateHeroNameReturnsErrorForShortName(t *testing.T) {
	err := ValidateHeroName("Ra")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "hero name is too short" {
		t.Fatalf("expected short name error, got %q", err.Error())
	}
}

func TestValidateHeroNameReturnsErrorForLongName(t *testing.T) {
	err := ValidateHeroName("VeryLongHeroNameOverLimit")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "hero name is too long" {
		t.Fatalf("expected long name error, got %q", err.Error())
	}
}

func TestIsValidHeroNameReturnsTrueForValidName(t *testing.T) {
	if !IsValidHeroName("Ragnar") {
		t.Fatal("expected hero name to be valid")
	}
}

func TestIsValidHeroNameReturnsFalseForInvalidName(t *testing.T) {
	if IsValidHeroName("Ra") {
		t.Fatal("expected hero name to be invalid")
	}
}

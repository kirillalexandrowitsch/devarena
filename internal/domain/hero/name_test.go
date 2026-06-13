package hero

import (
	"errors"
	"testing"
)

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

func TestValidateHeroNameReturnsErrHeroNameEmpty(t *testing.T) {
	err := ValidateHeroName("")

	if !errors.Is(err, ErrHeroNameEmpty) {
		t.Fatalf("expected ErrHeroNameEmpty, got %v", err)
	}
}

func TestValidateHeroNameReturnsValidationErrorForShortName(t *testing.T) {
	err := ValidateHeroName("Ra")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected ValidationError, got %T", err)
	}

	if validationErr.Field != "hero_name" {
		t.Fatalf("expected field %q, got %q", "hero_name", validationErr.Field)
	}

	expectedMessage := `value "Ra" is too short: length 2 is less than 3`
	if validationErr.Message != expectedMessage {
		t.Fatalf("expected message %q, got %q", expectedMessage, validationErr.Message)
	}
}

func TestValidateHeroNameReturnsValidationErrorForLongName(t *testing.T) {
	err := ValidateHeroName("VeryLongHeroNameOverLimit")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected ValidationError, got %T", err)
	}

	if validationErr.Field != "hero_name" {
		t.Fatalf("expected field %q, got %q", "hero_name", validationErr.Field)
	}

	expectedMessage := `value "VeryLongHeroNameOverLimit" is too long: length 25 is greater than 20`
	if validationErr.Message != expectedMessage {
		t.Fatalf("expected message %q, got %q", expectedMessage, validationErr.Message)
	}
}

func TestIsValidHeroNameReturnsTrueForValidName(t *testing.T) {
	if !IsValidHeroName("Ragnar") {
		t.Fatal("expected hero name to be valid")
	}
}

func TestIsValidHeroNameReturnsFalseForInvalidName(t *testing.T) {
	if IsValidHeroName("") {
		t.Fatal("expected hero name to be invalid")
	}
}

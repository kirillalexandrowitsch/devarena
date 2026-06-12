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

func TestValidateHeroNameReturnsErrHeroNameEmpty(t *testing.T) {
	err := ValidateHeroName("")

	if err != ErrHeroNameEmpty {
		t.Fatalf("expected ErrHeroNameEmpty, got %v", err)
	}
}

func TestValidateHeroNameReturnsFormattedErrorForShortName(t *testing.T) {
	err := ValidateHeroName("Ra")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	expectedMessage := `hero name "Ra" is too short: length 2 is less than 3`
	if err.Error() != expectedMessage {
		t.Fatalf("expected error %q, got %q", expectedMessage, err.Error())
	}
}

func TestValidateHeroNameReturnsFormattedErrorForLongName(t *testing.T) {
	err := ValidateHeroName("VeryLongHeroNameOverLimit")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	expectedMessage := `hero name "VeryLongHeroNameOverLimit" is too long: length 25 is greater than 20`
	if err.Error() != expectedMessage {
		t.Fatalf("expected error %q, got %q", expectedMessage, err.Error())
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

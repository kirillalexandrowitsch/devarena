package hero

import "testing"

func TestHeroNameLengthCountsRunes(t *testing.T) {
	name := "Ragnar Герой"

	got := HeroNameLength(name)
	want := 12

	if got != want {
		t.Fatalf("expected hero name length %d, got %d", want, got)
	}
}

func TestIsValidHeroName(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "Rag", want: true},
		{name: "Герой", want: true},
		{name: "Ra", want: false},
		{name: "VeryVeryVeryLongHeroName", want: false},
	}

	for _, tt := range tests {
		got := IsValidHeroName(tt.name)
		if got != tt.want {
			t.Fatalf("expected IsValidHeroName(%q) to be %v, got %v", tt.name, tt.want, got)
		}
	}
}

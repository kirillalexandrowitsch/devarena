package hero

import "testing"

func TestNewWeaponSnapshotReturnsNotEquippedForNilWeapon(t *testing.T) {
	snapshot := NewWeaponSnapshot(nil)

	if snapshot.Equipped {
		t.Fatal("expected weapon not to be equipped")
	}
}

func TestNewWeaponSnapshotReturnsSwordDetails(t *testing.T) {
	weapon := Sword{
		Title: "Starter Sword",
		Bonus: 4,
	}

	snapshot := NewWeaponSnapshot(weapon)

	if !snapshot.Equipped {
		t.Fatal("expected weapon to be equipped")
	}

	if snapshot.Name != "Starter Sword" {
		t.Fatalf("expected weapon name %q, got %q", "Starter Sword", snapshot.Name)
	}

	if snapshot.DamageBonus != 4 {
		t.Fatalf("expected weapon bonus %d, got %d", 4, snapshot.DamageBonus)
	}

	if snapshot.DynamicType != "hero.Sword" {
		t.Fatalf("expected dynamic type %q, got %q", "hero.Sword", snapshot.DynamicType)
	}
}

func TestNewWeaponSnapshotReturnsAxeDetails(t *testing.T) {
	weapon := Axe{
		Title: "Training Axe",
		Bonus: 6,
	}

	snapshot := NewWeaponSnapshot(weapon)

	if !snapshot.Equipped {
		t.Fatal("expected weapon to be equipped")
	}

	if snapshot.Name != "Training Axe" {
		t.Fatalf("expected weapon name %q, got %q", "Training Axe", snapshot.Name)
	}

	if snapshot.DamageBonus != 6 {
		t.Fatalf("expected weapon bonus %d, got %d", 6, snapshot.DamageBonus)
	}

	if snapshot.DynamicType != "hero.Axe" {
		t.Fatalf("expected dynamic type %q, got %q", "hero.Axe", snapshot.DynamicType)
	}
}

func TestHeroWeaponSnapshotReturnsEquippedWeapon(t *testing.T) {
	gameHero := Hero{
		Weapon: Sword{
			Title: "Starter Sword",
			Bonus: 4,
		},
	}

	snapshot := gameHero.WeaponSnapshot()

	if !snapshot.Equipped {
		t.Fatal("expected hero weapon to be equipped")
	}

	if snapshot.Name != "Starter Sword" {
		t.Fatalf("expected weapon name %q, got %q", "Starter Sword", snapshot.Name)
	}
}

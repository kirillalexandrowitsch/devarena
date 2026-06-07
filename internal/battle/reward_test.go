package battle

import "testing"

func TestCalculateRewardForGoblin(t *testing.T) {
	reward := CalculateReward(1, "Goblin")

	if reward.Experience != 30 {
		t.Fatalf("expected experience %d, got %d", 30, reward.Experience)
	}

	if reward.Item != "Goblin Dagger" {
		t.Fatalf("expected item %q, got %q", "Goblin Dagger", reward.Item)
	}
}

func TestCalculateRewardForUnknownEnemy(t *testing.T) {
	reward := CalculateReward(1, "Unknown Enemy")

	if reward.Experience != 30 {
		t.Fatalf("expected experience %d, got %d", 30, reward.Experience)
	}

	if reward.Item != "Rusty Sword" {
		t.Fatalf("expected item %q, got %q", "Rusty Sword", reward.Item)
	}
}

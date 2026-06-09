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

func TestEnemyRewardItemRule(t *testing.T) {
	rule := newEnemyRewardItemRule("Goblin", "Goblin Dagger")

	item, matched := rule("Goblin")

	if !matched {
		t.Fatal("expected reward rule to match enemy")
	}

	if item != "Goblin Dagger" {
		t.Fatalf("expected item %q, got %q", "Goblin Dagger", item)
	}
}

func TestEnemyRewardItemRuleDoesNotMatchDifferentEnemy(t *testing.T) {
	rule := newEnemyRewardItemRule("Goblin", "Goblin Dagger")

	item, matched := rule("Orc")

	if matched {
		t.Fatal("expected reward rule not to match enemy")
	}

	if item != "" {
		t.Fatalf("expected empty item, got %q", item)
	}
}

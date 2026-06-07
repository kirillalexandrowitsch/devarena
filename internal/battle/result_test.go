package battle

import (
	"testing"

	"github.com/rudyakovk/devarena/internal/enemy"
	"github.com/rudyakovk/devarena/internal/hero"
)

func TestBattleResultNamesWhenEnemyDefeated(t *testing.T) {
	gameHero := hero.Hero{
		Name: "Ragnar",
	}

	gameEnemy := enemy.Enemy{
		Name: "Goblin",
		HP:   0,
	}

	gameBattle := Battle{
		Hero:  &gameHero,
		Enemy: &gameEnemy,
	}

	winnerName, defeatedName := gameBattle.ResultNames()

	if winnerName != "Ragnar" {
		t.Fatalf("expected winner %q, got %q", "Ragnar", winnerName)
	}

	if defeatedName != "Goblin" {
		t.Fatalf("expected defeated %q, got %q", "Goblin", defeatedName)
	}
}

func TestBattleResultNamesWhenHeroDefeated(t *testing.T) {
	gameHero := hero.Hero{
		Name: "Ragnar",
		CombatStats: hero.CombatStats{
			HP: 0,
		},
	}

	gameEnemy := enemy.Enemy{
		Name: "Goblin",
		HP:   10,
	}

	gameBattle := Battle{
		Hero:  &gameHero,
		Enemy: &gameEnemy,
	}

	winnerName, defeatedName := gameBattle.ResultNames()

	if winnerName != "Goblin" {
		t.Fatalf("expected winner %q, got %q", "Goblin", winnerName)
	}

	if defeatedName != "Ragnar" {
		t.Fatalf("expected defeated %q, got %q", "Ragnar", defeatedName)
	}
}

func TestBattleResultNamesWhenBattleIsNotFinished(t *testing.T) {
	gameHero := hero.Hero{
		Name: "Ragnar",
		CombatStats: hero.CombatStats{
			HP: 100,
		},
	}

	gameEnemy := enemy.Enemy{
		Name: "Goblin",
		HP:   10,
	}

	gameBattle := Battle{
		Hero:  &gameHero,
		Enemy: &gameEnemy,
	}

	winnerName, defeatedName := gameBattle.ResultNames()

	if winnerName != "" {
		t.Fatalf("expected empty winner, got %q", winnerName)
	}

	if defeatedName != "" {
		t.Fatalf("expected empty defeated name, got %q", defeatedName)
	}
}

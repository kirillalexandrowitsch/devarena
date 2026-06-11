package arena

import (
	"github.com/rudyakovk/devarena/internal/domain/battle"
	"github.com/rudyakovk/devarena/internal/domain/enemy"
	"github.com/rudyakovk/devarena/internal/domain/hero"
)

const (
	defaultHeroName           = "Ragnar"
	defaultHeroClass          = hero.HeroClassWarrior
	defaultHeroLevel          = 1
	defaultHeroHP             = 100
	defaultHeroBaseDamage     = 15
	defaultHeroBonusDamage    = 5
	defaultHeroCriticalChance = 0.15

	defaultEnemyName = "Goblin"
	defaultEnemyHP   = 60
)

func Run() {
	NewApp().Run()
}

func (app *App) Run() {
	sessionReport := startSessionReport()

	defer func() {
		sessionReport.finish()

		if app.printSessionReport {
			sessionReport.printTo(app.output)
		}
	}()

	app.println("Welcome to DevArena")

	gameHero := mustCreateDefaultHero()

	starterSword := hero.Sword{
		Title: "Starter Sword",
		Bonus: 4,
	}

	gameHero.EquipWeapon(starterSword)

	if app.showInventoryInfo {
		if gameHero.HasItem("Small Potion") {
			app.println("Hero has a Small Potion")
		}

		firstInventoryItem := gameHero.FirstInventoryItem()
		if firstInventoryItem.Found {
			app.println("First inventory item:", firstInventoryItem.Value)
		}

		inventorySnapshot := gameHero.InventorySnapshot()
		app.println("Hero inventory length:", inventorySnapshot.Length)
		app.println("Hero inventory capacity:", inventorySnapshot.Capacity)
	}

	strength, strengthExists := gameHero.Stat("strength")
	if strengthExists {
		gameHero.BonusDamage += strength / 5
	}

	temporaryBonus, temporaryBonusExists := gameHero.Stat("temporary_bonus")
	if temporaryBonusExists {
		gameHero.BonusDamage += temporaryBonus
		gameHero.RemoveStat("temporary_bonus")
	}

	if app.showStatSummary {
		highestStatValue := gameHero.HighestStatValue()
		if highestStatValue.Found {
			app.println("Highest hero stat value:", highestStatValue.Value)
		}

		statsSnapshot := gameHero.StatsSnapshot()
		app.println("Hero stats count:", statsSnapshot.Count)
	}

	app.println("Hero class description:", hero.DescribeHeroClass(gameHero.Class))

	gameHero.PrintInfo()
	gameHero.PrintStats()
	gameHero.PrintAttacks()
	gameHero.PrintInventory("Hero inventory before battle:")

	gameEnemy := enemy.Enemy{
		Name: defaultEnemyName,
		HP:   defaultEnemyHP,
	}

	arenaBattle := battle.Battle{
		Hero:  &gameHero,
		Enemy: &gameEnemy,
		Round: 0,
	}

	arenaBattle.Start()

	winnerName, defeatedName := arenaBattle.ResultNames()

	battleStatistics := battle.Statistics{
		Wins:   1,
		Losses: 0,
	}

	reward := battle.CalculateReward(gameHero.Level, defeatedName)
	gameHero.AddItem(reward.Item)

	rewardEvent := battle.NewRewardEvent(gameHero.Name, reward)

	app.println(gameHero.Name, "received experience:", reward.Experience)
	app.println(gameHero.Name, "received item:", reward.Item)

	app.println("Reward event type:", rewardEvent.Type)
	app.println("Reward event payload size:", rewardEvent.Payload.Size())

	rewardItemMetadata, rewardItemMetadataExists := rewardEvent.MetadataString("reward_item")
	if rewardItemMetadataExists {
		app.println("Reward event metadata item:", rewardItemMetadata)
	}

	rewardExperienceMetadata, rewardExperienceMetadataExists := rewardEvent.MetadataInt("reward_experience")
	if rewardExperienceMetadataExists {
		app.println("Reward event metadata experience:", rewardExperienceMetadata)
	}

	rewardGrantedMetadata, rewardGrantedMetadataExists := rewardEvent.MetadataText("reward_granted")
	if rewardGrantedMetadataExists {
		app.println("Reward event metadata granted:", rewardGrantedMetadata)
	}

	app.println("Battle winner:", winnerName)
	app.println("Defeated opponent:", defeatedName)

	app.println("Battle win rate:", battleStatistics.WinRatePercent())

	gameHero.PrintInventory("Hero inventory after battle:")

	app.println("Final hero HP:", gameHero.HP)
	app.println("Final enemy HP:", gameEnemy.HP)
}

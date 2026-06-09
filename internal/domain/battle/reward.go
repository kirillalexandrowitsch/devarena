package battle

type Reward struct {
	Experience int
	Item       string
}

type RewardItemRule func(defeatedEnemyName string) (string, bool)

func CalculateReward(heroLevel int, defeatedEnemyName string) Reward {
	baseExperience := 25
	levelBonus := heroLevel * 5

	rewardCandidates := []string{
		"Rusty Sword",
	}

	rewardRules := []RewardItemRule{
		newEnemyRewardItemRule("Goblin", "Goblin Dagger"),
	}

	for _, rule := range rewardRules {
		item, matched := rule(defeatedEnemyName)
		if !matched {
			continue
		}

		rewardCandidates = []string{
			item,
			"Rusty Sword",
		}
		break
	}

	return Reward{
		Experience: baseExperience + levelBonus,
		Item:       selectRewardItem(rewardCandidates),
	}
}

func newEnemyRewardItemRule(enemyName string, rewardItem string) RewardItemRule {
	return func(defeatedEnemyName string) (string, bool) {
		if defeatedEnemyName != enemyName {
			return "", false
		}

		return rewardItem, true
	}
}

func selectRewardItem(candidates []string) string {
	selectedReward := "Rusty Sword"

	for _, item := range candidates {
		if item == "" {
			continue
		}

		selectedReward = item
		break
	}

	return selectedReward
}

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
		func(enemyName string) (string, bool) {
			if enemyName == "Goblin" {
				return "Goblin Dagger", true
			}

			return "", false
		},
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

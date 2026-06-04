package battle

type Reward struct {
	Experience int
	Item       string
}

func CalculateReward(heroLevel int, defeatedEnemyName string) Reward {
	baseExperience := 25
	levelBonus := heroLevel * 5

	rewardCandidates := []string{
		"Rusty Sword",
	}

	if defeatedEnemyName == "Goblin" {
		rewardCandidates = []string{
			"Goblin Dagger",
			"Rusty Sword",
		}
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

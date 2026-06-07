package battle

func (b Battle) ResultNames() (winnerName string, defeatedName string) {
	if b.Enemy.HP <= 0 {
		winnerName = b.Hero.Name
		defeatedName = b.Enemy.Name

		return
	}

	if b.Hero.HP <= 0 {
		winnerName = b.Enemy.Name
		defeatedName = b.Hero.Name

		return
	}

	return
}

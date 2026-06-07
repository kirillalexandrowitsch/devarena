package battle

func (b Battle) ResultNames() (string, string) {
	if b.Enemy.HP <= 0 {
		return b.Hero.Name, b.Enemy.Name
	}

	if b.Hero.HP <= 0 {
		return b.Enemy.Name, b.Hero.Name
	}

	return "", ""
}

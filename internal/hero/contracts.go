package hero

type DamageCalculator interface {
	TotalDamage() int
}

type InventoryManager interface {
	AddItem(item string)
}

type BattleRewardReceiver interface {
	DamageCalculator
	InventoryManager
}

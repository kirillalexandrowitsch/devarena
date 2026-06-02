package hero

type HeroClass string

const (
	HeroClassWarrior HeroClass = "Warrior"
	HeroClassMage    HeroClass = "Mage"
	HeroClassRogue   HeroClass = "Rogue"
)

func DescribeHeroClass(class HeroClass) string {
	switch class {
	case HeroClassWarrior:
		return "Warrior is a durable melee fighter with stable damage"
	case HeroClassMage:
		return "Mage is a ranged fighter with high magic damage"
	case HeroClassRogue:
		return "Rogue is a fast fighter with critical attacks"
	default:
		return "Unknown hero class"
	}
}

package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/rudyakovk/devarena/internal/battle"
	"github.com/rudyakovk/devarena/internal/enemy"
	"github.com/rudyakovk/devarena/internal/hero"
)

const (
	defaultHeroClass          = hero.HeroClassWarrior
	defaultHeroLevel          = 1
	defaultHeroHP             = 100
	defaultHeroBaseDamage     = 15
	defaultHeroBonusDamage    = 5
	defaultHeroCriticalChance = 0.15

	defaultEnemyName = "Goblin"
	defaultEnemyHP   = 60
)

func damageHeroCopy(hero hero.Hero, damage int) {
	hero.HP -= damage
	fmt.Println("Inside damageHeroCopy HP:", hero.HP)
}

func damageHeroOriginal(hero *hero.Hero, damage int) {
	hero.HP -= damage

	if hero.HP <= 0 {
		hero.HP = 0
		hero.Alive = false
	}

	fmt.Println("Inside damageHeroOriginal HP:", hero.HP)
}

func printWeaponInfo(weapon hero.Weapon) {
	fmt.Println("Weapon name:", weapon.Name())
	fmt.Println("Weapon damage bonus:", weapon.DamageBonus())
}

func printCalculatedDamage(calculator hero.DamageCalculator) {
	fmt.Println("Calculated damage from interface:", calculator.TotalDamage())
}

func addMethodSetReward(manager hero.InventoryManager) {
	manager.AddItem("Method Set Badge")
}

func addComposedInterfaceReward(receiver hero.BattleRewardReceiver) {
	fmt.Println("Composed interface damage:", receiver.TotalDamage())
	receiver.AddItem("Interface Composition Medal")
}

func printAnyValue(label string, value any) {
	fmt.Println(label+":", value)
}

func printStringValue(label string, value any) {
	text, ok := value.(string)
	if !ok {
		fmt.Println(label + " is not a string")
		return
	}

	fmt.Println(label+" as string:", text)
}

func printTypedValue(label string, value any) {
	switch typedValue := value.(type) {
	case string:
		fmt.Println(label+" is string:", typedValue)
	case int:
		fmt.Println(label+" is int:", typedValue)
	case bool:
		fmt.Println(label+" is bool:", typedValue)
	case []string:
		fmt.Println(label+" is []string:", typedValue)
	case hero.Hero:
		fmt.Println(label+" is hero.Hero:", typedValue.Name)
	default:
		fmt.Println(label + " has unknown type")
	}
}

func printZeroValues() {
	var defaultName string
	var defaultLevel int
	var defaultCriticalChance float64
	var defaultAlive bool
	var defaultHeroPointer *hero.Hero
	var defaultInventory []string
	var defaultStats map[string]int
	var defaultHero hero.Hero

	fmt.Println("Zero values demo:")
	fmt.Println("Default string:", defaultName)
	fmt.Println("Default int:", defaultLevel)
	fmt.Println("Default float64:", defaultCriticalChance)
	fmt.Println("Default bool:", defaultAlive)
	fmt.Println("Default hero pointer:", defaultHeroPointer)
	fmt.Println("Default inventory slice:", defaultInventory)
	fmt.Println("Default stats map:", defaultStats)
	fmt.Println("Default hero:", defaultHero)
}

func printTypeConversions() {
	wins := 7
	losses := 3
	totalBattles := wins + losses

	winRate := float64(wins) / float64(totalBattles) * 100
	totalBattles64 := int64(totalBattles)
	winRateText := fmt.Sprintf("%.2f%%", winRate)

	fmt.Println("Type conversion demo:")
	fmt.Println("Wins:", wins)
	fmt.Println("Losses:", losses)
	fmt.Println("Total battles as int:", totalBattles)
	fmt.Println("Total battles as int64:", totalBattles64)
	fmt.Println("Win rate as float64:", winRate)
	fmt.Println("Win rate text:", winRateText)
}

func printRuneDemo() {
	heroTitle := "Ragnar Герой"
	titleRunes := []rune(heroTitle)

	fmt.Println("Rune demo:")
	fmt.Println("Hero title:", heroTitle)
	fmt.Println("Bytes length:", len(heroTitle))
	fmt.Println("Runes length:", len(titleRunes))

	for index, symbol := range heroTitle {
		fmt.Printf("Byte index %d: %c (%U)\n", index, symbol, symbol)
	}
}

func printByteDemo() {
	heroTitle := "Ragnar Герой"
	titleBytes := []byte(heroTitle)

	fmt.Println("Byte demo:")
	fmt.Println("Hero title:", heroTitle)
	fmt.Println("Bytes length:", len(titleBytes))
	fmt.Println("Bytes:", titleBytes)

	for index, symbolByte := range titleBytes {
		fmt.Printf("Byte index %d: %d (%X)\n", index, symbolByte, symbolByte)
	}
}

func prepareArena() {
	fmt.Println("Arena preparation started")

	defer fmt.Println("Arena preparation log closed")
	defer fmt.Println("Arena temporary resources released")

	fmt.Println("Checking arena gates")
	fmt.Println("Checking arena weapons")
	fmt.Println("Arena preparation finished")
}

func validateArenaCapacity(capacity int) {
	if capacity <= 0 {
		panic("arena capacity must be positive")
	}

	fmt.Println("Arena capacity is valid:", capacity)
}

func runArenaSafetyCheck(capacity int) {
	defer func() {
		recoveredValue := recover()
		if recoveredValue != nil {
			fmt.Println("Recovered from arena panic:", recoveredValue)
		}
	}()

	validateArenaCapacity(capacity)

	fmt.Println("Arena safety check completed")
}

func firstOrDefault[T any](items []T, defaultValue T) T {
	if len(items) == 0 {
		return defaultValue
	}

	return items[0]
}

func printGenericsDemo() {
	starterItems := []string{"Small Potion", "Wooden Shield"}
	enemyLevels := []int{1, 2, 3}
	arenaReadiness := []bool{true}

	firstItem := firstOrDefault(starterItems, "No item")
	firstEnemyLevel := firstOrDefault(enemyLevels, 0)
	firstReadinessState := firstOrDefault(arenaReadiness, false)
	emptyReward := firstOrDefault([]string{}, "Default Reward")

	fmt.Println("Generics demo:")
	fmt.Println("First starter item:", firstItem)
	fmt.Println("First enemy level:", firstEnemyLevel)
	fmt.Println("First readiness state:", firstReadinessState)
	fmt.Println("Empty reward fallback:", emptyReward)
}

func printPair[T any, U any](label string, first T, second U) {
	fmt.Println(label+":", first, "->", second)
}

func printTypeParametersDemo() {
	fmt.Println("Type parameters demo:")

	printPair("Hero level to reward", 1, "Goblin Dagger")
	printPair("Arena readiness to capacity", true, 100)
	printPair[string, int]("Hero name to level", "Ragnar", 1)
}

type arenaNumber interface {
	~int | ~int64 | ~float64
}

func maxArenaValue[T arenaNumber](first T, second T) T {
	if first > second {
		return first
	}

	return second
}

func printConstraintsDemo() {
	maxHeroLevel := maxArenaValue(1, 5)
	maxArenaCapacity := maxArenaValue[int64](100, 250)
	maxCriticalChance := maxArenaValue(0.15, 0.35)

	fmt.Println("Constraints demo:")
	fmt.Println("Max hero level:", maxHeroLevel)
	fmt.Println("Max arena capacity:", maxArenaCapacity)
	fmt.Println("Max critical chance:", maxCriticalChance)
}

type ArenaConfig struct {
	Name     string
	Capacity int
	Ranked   bool
}

type ArenaOption func(*ArenaConfig)

func WithArenaName(name string) ArenaOption {
	return func(config *ArenaConfig) {
		config.Name = name
	}
}

func WithArenaCapacity(capacity int) ArenaOption {
	return func(config *ArenaConfig) {
		config.Capacity = capacity
	}
}

func WithRankedArena(ranked bool) ArenaOption {
	return func(config *ArenaConfig) {
		config.Ranked = ranked
	}
}

func NewArenaConfig(options ...ArenaOption) ArenaConfig {
	config := ArenaConfig{
		Name:     "Training Arena",
		Capacity: 100,
		Ranked:   false,
	}

	for _, option := range options {
		option(&config)
	}

	return config
}

func printFunctionalOptionsDemo() {
	defaultArena := NewArenaConfig()

	tournamentArena := NewArenaConfig(
		WithArenaName("Goblin Cup Arena"),
		WithArenaCapacity(250),
		WithRankedArena(true),
	)

	fmt.Println("Functional options demo:")
	fmt.Println("Default arena:", defaultArena)
	fmt.Println("Tournament arena:", tournamentArena)
}

type ArenaNotifier interface {
	Notify(message string)
}

type ConsoleNotifier struct{}

func (notifier ConsoleNotifier) Notify(message string) {
	fmt.Println("Arena notification:", message)
}

type ArenaService struct {
	notifier ArenaNotifier
}

func NewArenaService(notifier ArenaNotifier) ArenaService {
	return ArenaService{
		notifier: notifier,
	}
}

func (service ArenaService) AnnounceArenaReady() {
	service.notifier.Notify("Arena service is ready")
}

func printDependencyInjectionDemo() {
	consoleNotifier := ConsoleNotifier{}
	arenaService := NewArenaService(consoleNotifier)

	fmt.Println("Dependency injection demo:")
	arenaService.AnnounceArenaReady()
}

func printSliceInternalsDemo() {
	baseInventory := []string{"Small Potion", "Wooden Shield", "Old Boots"}
	inventoryView := baseInventory[:2]

	fmt.Println("Slice internals demo:")
	fmt.Println("Base inventory:", baseInventory)
	fmt.Println("Inventory view:", inventoryView)
	fmt.Println("Base inventory len:", len(baseInventory), "cap:", cap(baseInventory))
	fmt.Println("Inventory view len:", len(inventoryView), "cap:", cap(inventoryView))

	inventoryView[0] = "Large Potion"

	fmt.Println("Base inventory after view update:", baseInventory)
	fmt.Println("Inventory view after update:", inventoryView)

	extendedView := append(inventoryView, "Iron Helmet")

	fmt.Println("Extended view after append:", extendedView)
	fmt.Println("Base inventory after append to view:", baseInventory)

	isolatedInventory := append([]string{}, baseInventory...)
	isolatedInventory[0] = "Isolated Potion"

	fmt.Println("Base inventory after isolated update:", baseInventory)
	fmt.Println("Isolated inventory:", isolatedInventory)
}

func addInventoryCounter(counters map[string]int, item string, count int) {
	counters[item] = count
}

func printMapInternalsDemo() {
	inventoryCounters := map[string]int{
		"Small Potion":  2,
		"Wooden Shield": 1,
		"Iron Helmet":   1,
	}

	fmt.Println("Map internals demo:")
	fmt.Println("Inventory counters:", inventoryCounters)

	potionCount, potionExists := inventoryCounters["Small Potion"]
	fmt.Println("Small Potion count:", potionCount, "exists:", potionExists)

	unknownCount, unknownExists := inventoryCounters["Unknown Item"]
	fmt.Println("Unknown Item count:", unknownCount, "exists:", unknownExists)

	sharedCounters := inventoryCounters
	sharedCounters["Small Potion"] = 5

	fmt.Println("Inventory counters after shared map update:", inventoryCounters)
	fmt.Println("Shared counters:", sharedCounters)

	addInventoryCounter(inventoryCounters, "Goblin Dagger", 1)

	fmt.Println("Inventory counters after function update:", inventoryCounters)

	delete(inventoryCounters, "Wooden Shield")

	fmt.Println("Inventory counters after delete:", inventoryCounters)

	fmt.Println("Map iteration order:")
	for item, count := range inventoryCounters {
		fmt.Println(item, "->", count)
	}

	var nilCounters map[string]int

	fmt.Println("Nil map:", nilCounters)
	fmt.Println("Nil map length:", len(nilCounters))
	fmt.Println("Read from nil map:", nilCounters["Missing Item"])

	initializedCounters := make(map[string]int)
	initializedCounters["Training Token"] = 1

	fmt.Println("Initialized map:", initializedCounters)
}

func printInterfaceInternalsDemo() {
	var emptyValue any

	var nilHeroPointer *hero.Hero
	var typedNilAsAny any = nilHeroPointer

	interfaceHero := hero.Hero{
		ID:    2,
		Name:  "Interface Hero",
		Class: hero.HeroClassWarrior,
		Level: 1,
		Alive: true,
		CombatStats: hero.CombatStats{
			HP:             100,
			BaseDamage:     10,
			BonusDamage:    2,
			CriticalChance: 0.10,
		},
	}

	var damageCalculator hero.DamageCalculator = interfaceHero
	var inventoryManager hero.InventoryManager = &interfaceHero

	fmt.Println("Interface internals demo:")
	fmt.Printf("Empty interface: type=%T value=%v is nil=%v\n", emptyValue, emptyValue, emptyValue == nil)
	fmt.Printf("Typed nil pointer as any: type=%T value=%v is nil=%v\n", typedNilAsAny, typedNilAsAny, typedNilAsAny == nil)
	fmt.Printf("Damage calculator dynamic type: %T\n", damageCalculator)
	fmt.Println("Damage calculator total damage:", damageCalculator.TotalDamage())
	fmt.Printf("Inventory manager dynamic type: %T\n", inventoryManager)

	inventoryManager.AddItem("Interface Internals Badge")

	fmt.Println("Hero inventory after interface call:", interfaceHero.Inventory)
}

func createArenaNameValue() string {
	arenaName := "Stack Training Arena"

	return arenaName
}

func createArenaNamePointer() *string {
	arenaName := "Heap Candidate Arena"

	return &arenaName
}

func printStackHeapDemo() {
	localScore := 100
	arenaNameValue := createArenaNameValue()
	arenaNamePointer := createArenaNamePointer()

	fmt.Println("Stack vs heap demo:")
	fmt.Println("Local score:", localScore)
	fmt.Printf("Local score address: %p\n", &localScore)
	fmt.Println("Arena name value:", arenaNameValue)
	fmt.Println("Arena name pointer value:", *arenaNamePointer)
	fmt.Printf("Arena name pointer address: %p\n", arenaNamePointer)
}

func createArenaScoreValue() int {
	score := 100

	return score
}

func createArenaScorePointer() *int {
	score := 100

	return &score
}

func printEscapeAnalysisDemo() {
	scoreValue := createArenaScoreValue()
	scorePointer := createArenaScorePointer()

	fmt.Println("Escape analysis demo:")
	fmt.Println("Score value:", scoreValue)
	fmt.Println("Score pointer value:", *scorePointer)
	fmt.Printf("Score pointer address: %p\n", scorePointer)
}

func printGarbageCollectorDemo() {
	var before runtime.MemStats
	var afterAllocation runtime.MemStats
	var afterGC runtime.MemStats

	runtime.ReadMemStats(&before)

	temporaryInventories := make([][]string, 0, 10000)

	for i := 0; i < 10000; i++ {
		inventory := []string{
			"Small Potion",
			"Wooden Shield",
			"Goblin Dagger",
		}

		temporaryInventories = append(temporaryInventories, inventory)
	}

	runtime.ReadMemStats(&afterAllocation)

	fmt.Println("Garbage collector demo:")
	fmt.Println("Temporary inventories count:", len(temporaryInventories))
	fmt.Println("Heap allocation before:", before.HeapAlloc)
	fmt.Println("Heap allocation after allocation:", afterAllocation.HeapAlloc)
	fmt.Println("GC cycles before:", before.NumGC)
	fmt.Println("GC cycles after allocation:", afterAllocation.NumGC)

	temporaryInventories = nil

	runtime.GC()
	runtime.ReadMemStats(&afterGC)

	fmt.Println("Heap allocation after manual GC:", afterGC.HeapAlloc)
	fmt.Println("GC cycles after manual GC:", afterGC.NumGC)
}

type ArenaMemoryStats struct {
	Active   bool
	Capacity int64
	Level    int32
	Ranked   bool
}

type OptimizedArenaMemoryStats struct {
	Capacity int64
	Level    int32
	Active   bool
	Ranked   bool
}

func printStructAlignmentDemo() {
	arenaStats := ArenaMemoryStats{}
	optimizedStats := OptimizedArenaMemoryStats{}

	fmt.Println("Struct alignment demo:")
	fmt.Println("ArenaMemoryStats size:", unsafe.Sizeof(arenaStats))
	fmt.Println("ArenaMemoryStats align:", unsafe.Alignof(arenaStats))
	fmt.Println("ArenaMemoryStats Active offset:", unsafe.Offsetof(arenaStats.Active))
	fmt.Println("ArenaMemoryStats Capacity offset:", unsafe.Offsetof(arenaStats.Capacity))
	fmt.Println("ArenaMemoryStats Level offset:", unsafe.Offsetof(arenaStats.Level))
	fmt.Println("ArenaMemoryStats Ranked offset:", unsafe.Offsetof(arenaStats.Ranked))

	fmt.Println("OptimizedArenaMemoryStats size:", unsafe.Sizeof(optimizedStats))
	fmt.Println("OptimizedArenaMemoryStats align:", unsafe.Alignof(optimizedStats))
	fmt.Println("OptimizedArenaMemoryStats Capacity offset:", unsafe.Offsetof(optimizedStats.Capacity))
	fmt.Println("OptimizedArenaMemoryStats Level offset:", unsafe.Offsetof(optimizedStats.Level))
	fmt.Println("OptimizedArenaMemoryStats Active offset:", unsafe.Offsetof(optimizedStats.Active))
	fmt.Println("OptimizedArenaMemoryStats Ranked offset:", unsafe.Offsetof(optimizedStats.Ranked))
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

func calculateBattleReward(heroLevel int, defeatedEnemyName string) (rewardExperience int, rewardItem string) {
	baseExperience := 25
	levelBonus := heroLevel * 5

	rewardExperience = baseExperience + levelBonus

	rewardCandidates := []string{
		"",
		"Rusty Sword",
	}

	if defeatedEnemyName == "Goblin" {
		rewardCandidates = []string{
			"",
			"Goblin Dagger",
			"Rusty Sword",
		}
	}

	rewardItem = selectRewardItem(rewardCandidates)

	return
}

func main() {
	fmt.Println("Welcome to DevArena")

	printZeroValues()

	printTypeConversions()

	printRuneDemo()

	printByteDemo()

	prepareArena()

	validateArenaCapacity(100)
	runArenaSafetyCheck(0)

	printGenericsDemo()

	printTypeParametersDemo()

	printConstraintsDemo()

	printFunctionalOptionsDemo()

	printDependencyInjectionDemo()

	printSliceInternalsDemo()

	printMapInternalsDemo()

	printInterfaceInternalsDemo()

	printStackHeapDemo()

	printEscapeAnalysisDemo()

	printGarbageCollectorDemo()

	printStructAlignmentDemo()

	gameHero := hero.Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: defaultHeroClass,
		Level: defaultHeroLevel,
		Alive: true,
		CombatStats: hero.CombatStats{
			HP:             defaultHeroHP,
			BaseDamage:     defaultHeroBaseDamage,
			BonusDamage:    defaultHeroBonusDamage,
			CriticalChance: defaultHeroCriticalChance,
		},
		Attacks:   [3]string{"Slash", "Pierce", "Heavy Strike"},
		Inventory: []string{"Small Potion", "Wooden Shield"},
		Stats: map[string]int{
			"strength": 10,
			"agility":  7,
			"stamina":  12,
		},
	}

	fmt.Println("Hero class description:", hero.DescribeHeroClass(gameHero.Class))

	starterSword := hero.Sword{
		Title: "Starter Sword",
		Bonus: 4,
	}

	battleAxe := hero.Axe{
		Title: "Battle Axe",
		Bonus: 7,
	}

	fmt.Println("Available weapons:")
	printWeaponInfo(starterSword)
	printWeaponInfo(battleAxe)

	gameHero.EquipWeapon(starterSword)

	trainingSessions := 0

	calculateTrainingBonus := func(strength int) int {
		trainingSessions++

		return strength/2 + trainingSessions
	}

	firstTrainingBonus := calculateTrainingBonus(gameHero.Stats["strength"])
	secondTrainingBonus := calculateTrainingBonus(gameHero.Stats["strength"])
	trainingBonusDamage := firstTrainingBonus + secondTrainingBonus

	gameHero.BonusDamage += trainingBonusDamage

	fmt.Println("Training sessions:", trainingSessions)
	fmt.Println("First training bonus damage:", firstTrainingBonus)
	fmt.Println("Second training bonus damage:", secondTrainingBonus)
	fmt.Println("Total training bonus damage:", trainingBonusDamage)

	gameEnemy := enemy.Enemy{
		Name: defaultEnemyName,
		HP:   defaultEnemyHP,
	}

	gameBattle := battle.Battle{
		Hero:  &gameHero,
		Enemy: &gameEnemy,
		Round: 0,
	}

	inventoryBeforeBattle := make([]string, len(gameHero.Inventory))
	copy(inventoryBeforeBattle, gameHero.Inventory)

	gameHero.PrintInfo()
	gameHero.PrintStats()

	fmt.Println("Method set demo:")
	printCalculatedDamage(gameHero)
	printCalculatedDamage(&gameHero)
	addMethodSetReward(&gameHero)

	fmt.Println("Interface composition demo:")
	addComposedInterfaceReward(&gameHero)

	fmt.Println("Any demo:")
	printAnyValue("Hero name", gameHero.Name)
	printAnyValue("Hero level", gameHero.Level)
	printAnyValue("Hero alive", gameHero.Alive)
	printAnyValue("Hero inventory", gameHero.Inventory)
	printAnyValue("Hero struct", gameHero)

	fmt.Println("Type assertion demo:")
	printStringValue("Hero name", gameHero.Name)
	printStringValue("Hero level", gameHero.Level)

	fmt.Println("Type switch demo:")
	printTypedValue("Hero name", gameHero.Name)
	printTypedValue("Hero level", gameHero.Level)
	printTypedValue("Hero alive", gameHero.Alive)
	printTypedValue("Hero inventory", gameHero.Inventory)
	printTypedValue("Hero struct", gameHero)
	printTypedValue("Hero critical chance", gameHero.CriticalChance)

	intellect, exists := gameHero.Stats["intellect"]
	if exists {
		fmt.Println("Intellect:", intellect)
	} else {
		fmt.Println("Intellect stat is not defined")
	}

	gameHero.Stats["strength"] = gameHero.Stats["strength"] + 2
	gameHero.Stats["intellect"] = 3

	fmt.Println("Hero stats after training:")
	gameHero.PrintStats()

	delete(gameHero.Stats, "intellect")

	fmt.Println("Hero stats after removing temporary intellect bonus:")
	gameHero.PrintStats()

	gameHero.PrintAttacks()

	fmt.Println("Hero inventory before battle:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length before reward:", len(gameHero.Inventory))
	fmt.Println("Inventory capacity before reward:", cap(gameHero.Inventory))

	gameHero.AddItems("Rusty Sword", "Training Token")

	gameHero.PrintInventory("Hero inventory after reward:")

	fmt.Println("Inventory copy before battle still:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length after reward:", len(gameHero.Inventory))
	fmt.Println("Inventory capacity after reward:", cap(gameHero.Inventory))

	if gameHero.HP > 0 {
		fmt.Println(gameHero.Name, "is ready to fight")
	} else {
		fmt.Println(gameHero.Name, "is defeated and cannot fight")
	}

	fmt.Println("Pointer demo:")
	fmt.Println("Hero HP before damageHeroCopy:", gameHero.HP)
	damageHeroCopy(gameHero, 10)
	fmt.Println("Hero HP after damageHeroCopy:", gameHero.HP)

	fmt.Println("Hero HP before damageHeroOriginal:", gameHero.HP)
	damageHeroOriginal(&gameHero, 10)
	fmt.Println("Hero HP after damageHeroOriginal:", gameHero.HP)

	fmt.Println("Enemy name:", gameEnemy.Name)
	fmt.Println("Enemy HP:", gameEnemy.HP)

	gameBattle.Start()

	rewardExperience, rewardItem := calculateBattleReward(gameHero.Level, gameEnemy.Name)

	gameHero.AddItems(rewardItem)

	fmt.Println(gameHero.Name, "received experience:", rewardExperience)
	fmt.Println(gameHero.Name, "received item:", rewardItem)
	fmt.Println("Final hero HP:", gameHero.HP)
	fmt.Println("Final enemy HP:", gameEnemy.HP)
}

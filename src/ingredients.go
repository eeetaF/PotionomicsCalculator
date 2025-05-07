package src

type Ingredient struct {
	Magimints [5]uint16
	Category  IngredientCategory
	Name      string
	Price     int
	Traits    []TraitStruct
}

type IngredientSearch struct {
	Magimints [5]uint16
	Name      string
	Traits    []TraitStruct
}

var Ingredients = []Ingredient{
	// Slime
	{[5]uint16{0, 0, 6, 0, 0}, Slime, "Sack of Slime", 7, []TraitStruct{}},
	{[5]uint16{3, 3, 3, 0, 0}, Slime, "Cubic Ooze", 16, []TraitStruct{}},
	{[5]uint16{0, 0, 18, 0, 0}, Slime, "Sack of Hive Slime", 21, []TraitStruct{}},
	{[5]uint16{30, 0, 0, 0, 0}, Slime, "Antlered Jelly", 28, []TraitStruct{
		{Sensation, false}, {Visual, true}}},
	{[5]uint16{0, 0, 30, 0, 0}, Slime, "Sack of Composite Slime", 36, []TraitStruct{}},
	{[5]uint16{9, 9, 12, 12, 0}, Slime, "Bubble Ooze", 60, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 48, 0}, Slime, "Feathered Gelatin", 62, []TraitStruct{
		{Sensation, true}, {Aroma, false}}},
	{[5]uint16{32, 0, 0, 32, 0}, Slime, "Shelled Pudding", 90, []TraitStruct{}},
	{[5]uint16{15, 15, 15, 15, 0}, Slime, "Copper Dollop", 95, []TraitStruct{}},
	{[5]uint16{0, 0, 44, 0, 22}, Slime, "Bedazzled Custard", 95, []TraitStruct{
		{Visual, true}}},
	{[5]uint16{0, 0, 0, 66, 0}, Slime, "Winged Gelatin", 103, []TraitStruct{
		{Sensation, true}, {Aroma, false}}},
	{[5]uint16{24, 24, 24, 24, 0}, Slime, "Silver Dollop", 138, []TraitStruct{}},
	{[5]uint16{33, 33, 33, 33, 0}, Slime, "Gold Dollop", 176, []TraitStruct{}},

	// Plant
	{[5]uint16{0, 6, 0, 0, 0}, Plant, "Mandrake Root", 6, []TraitStruct{}},
	{[5]uint16{0, 27, 0, 0, 0}, Plant, "Bog Beet", 27, []TraitStruct{
		{Taste, false}, {Sound, true}}},
	{[5]uint16{0, 30, 0, 0, 0}, Plant, "Reef Radish", 32, []TraitStruct{
		{Taste, false}, {Sound, true}}},
	{[5]uint16{0, 30, 0, 0, 0}, Plant, "Mandragon Root", 34, []TraitStruct{}},
	{[5]uint16{0, 48, 0, 0, 0}, Plant, "Acid Rutabaga", 54, []TraitStruct{
		{Taste, false}, {Sound, true}}},
	{[5]uint16{0, 32, 0, 32, 0}, Plant, "Daredevil Pepper", 90, []TraitStruct{}},
	{[5]uint16{10, 0, 20, 0, 30}, Plant, "Mosquito Plant", 105, []TraitStruct{}},
	{[5]uint16{0, 44, 0, 44, 0}, Plant, "Widowmaker Pepper", 126, []TraitStruct{}},
	{[5]uint16{20, 20, 15, 0, 15}, Plant, "Squid Vine", 135, []TraitStruct{}},
	{[5]uint16{16, 0, 40, 0, 40}, Plant, "Acid Pitfall Plant", 145, []TraitStruct{}},
	{[5]uint16{24, 24, 24, 0, 24}, Plant, "Harpy's Snare", 150, []TraitStruct{}},
	{[5]uint16{22, 0, 55, 0, 55}, Plant, "Barracuda Plant", 172, []TraitStruct{}},

	// Flower
	{[5]uint16{4, 0, 0, 0, 0}, Flower, "Fairy Flower Bulb", 14, []TraitStruct{
		{Aroma, true}}},
	{[5]uint16{0, 0, 0, 12, 0}, Flower, "Wraith Orchid", 19, []TraitStruct{}},
	{[5]uint16{12, 0, 0, 0, 0}, Flower, "Fairy Flower Bud", 23, []TraitStruct{
		{Aroma, true}}},
	{[5]uint16{18, 0, 0, 6, 0}, Flower, "Ghostlight Bloom", 28, []TraitStruct{}},
	{[5]uint16{20, 0, 0, 0, 0}, Flower, "Fairy Flower Bloom", 35, []TraitStruct{
		{Aroma, true}}},
	{[5]uint16{16, 0, 0, 0, 0}, Flower, "Bramble-Rose", 45, []TraitStruct{
		{Sensation, true}, {Sound, true}}},
	{[5]uint16{22, 0, 0, 0, 0}, Flower, "Inverted Bramble-Rose", 46, []TraitStruct{
		{Sensation, true}, {Sound, true}}},
	{[5]uint16{40, 0, 0, 20, 0}, Flower, "Fire Flower", 55, []TraitStruct{
		{Aroma, false}}},
	{[5]uint16{24, 0, 0, 8, 0}, Flower, "Djinn Blossom", 68, []TraitStruct{
		{Taste, true}, {Aroma, true}}},
	{[5]uint16{8, 24, 24, 0, 0}, Flower, "Courtier's Orchid", 72, []TraitStruct{
		{Aroma, true}}},
	{[5]uint16{0, 16, 0, 48, 0}, Flower, "Watchdog Daisy", 83, []TraitStruct{}},
	{[5]uint16{11, 33, 0, 33, 0}, Flower, "Orchid of the Ice Princess", 116, []TraitStruct{
		{Aroma, true}}},

	// Fruit
	{[5]uint16{6, 0, 0, 0, 0}, Fruit, "Feyberry", 4, []TraitStruct{}},
	{[5]uint16{18, 0, 0, 0, 0}, Fruit, "Puckberry", 16, []TraitStruct{}},
	{[5]uint16{0, 18, 6, 0, 0}, Fruit, "Figment Pomme", 26, []TraitStruct{}},
	{[5]uint16{30, 0, 0, 0, 0}, Fruit, "Bogeyberry", 30, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 40, 0}, Fruit, "Saltwatermelon", 44, []TraitStruct{
		{Visual, false}}},
	{[5]uint16{0, 10, 30, 0, 0}, Fruit, "Phantom Pomme", 64, []TraitStruct{
		{Taste, true}, {Sound, false}}},
	{[5]uint16{0, 0, 0, 64, 0}, Fruit, "Rottermelon", 68, []TraitStruct{
		{Visual, false}}},
	{[5]uint16{0, 33, 11, 0, 0}, Fruit, "Nightmare Pomme", 72, []TraitStruct{
		{Visual, true}, {Sound, true}}},
	{[5]uint16{0, 24, 8, 0, 0}, Fruit, "Daydream Pomme", 75, []TraitStruct{
		{Visual, true}, {Sound, true}}},
	{[5]uint16{0, 16, 0, 0, 48}, Fruit, "Geode Citrus", 94, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 76, 0}, Fruit, "Slaughtermelon", 105, []TraitStruct{
		{Visual, false}}},
	{[5]uint16{0, 22, 0, 0, 66}, Fruit, "Dragonegg Citrus", 124, []TraitStruct{}},
	{[5]uint16{48, 0, 48, 24, 24}, Fruit, "Charredonnay", 260, []TraitStruct{
		{Taste, false}}},

	// Fungus
	{[5]uint16{0, 4, 0, 0, 0}, Fungus, "Impstool Mushroom", 17, []TraitStruct{
		{Sensation, true}}},
	{[5]uint16{0, 12, 0, 0, 0}, Fungus, "Trollstool Mushroom", 20, []TraitStruct{
		{Sensation, true}}},
	{[5]uint16{0, 18, 0, 6, 0}, Fungus, "Miasma Spore", 30, []TraitStruct{}},
	{[5]uint16{0, 0, 30, 0, 0}, Fungus, "Hallucinatory Shroom", 36, []TraitStruct{
		{Taste, true}, {Sound, false}}},
	{[5]uint16{0, 20, 0, 0, 0}, Fungus, "Giantstool Mushroom", 40, []TraitStruct{
		{Sensation, true}}},
	{[5]uint16{0, 0, 48, 0, 0}, Fungus, "Delirium Shroom", 63, []TraitStruct{
		{Taste, true}, {Sound, false}}},
	{[5]uint16{16, 0, 0, 0, 48}, Fungus, "Creeping Mildew", 92, []TraitStruct{}},
	{[5]uint16{0, 48, 0, 16, 0}, Fungus, "Medusa Spore", 94, []TraitStruct{
		{Taste, false}, {Aroma, true}}},
	{[5]uint16{32, 64, 64, 32, 0}, Fungus, "Shallow Grave Enoki", 200, []TraitStruct{
		{Sensation, false}, {Aroma, false}}},

	// Bug
	{[5]uint16{0, 0, 4, 0, 0}, Bug, "Rotfly Larva", 10, []TraitStruct{
		{Taste, true}}},
	{[5]uint16{0, 0, 12, 0, 0}, Bug, "Rotfly Cocoon", 25, []TraitStruct{
		{Taste, true}}},
	{[5]uint16{12, 6, 0, 0, 0}, Bug, "Sphinx Flea", 35, []TraitStruct{
		{Sensation, true}}},
	{[5]uint16{0, 0, 20, 0, 0}, Bug, "Rotfly Adult", 38, []TraitStruct{
		{Taste, true}}},
	{[5]uint16{10, 20, 0, 0, 0}, Bug, "Selkie Lice", 50, []TraitStruct{
		{Sensation, true}}},
	{[5]uint16{0, 0, 0, 0, 30}, Bug, "Static Spiderling", 50, []TraitStruct{
		{Visual, false}, {Sound, true}}},
	{[5]uint16{0, 0, 32, 0, 0}, Bug, "Rotfly Matriarch", 65, []TraitStruct{
		{Taste, true}}},
	{[5]uint16{0, 0, 44, 0, 0}, Bug, "Rotfly Mutant", 74, []TraitStruct{
		{Taste, true}}},
	{[5]uint16{0, 0, 0, 0, 48}, Bug, "Sepulcher Widow", 82, []TraitStruct{
		{Visual, false}, {Sound, true}}},
	{[5]uint16{0, 24, 24, 24, 0}, Bug, "Jeweled Scarab", 105, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 0, 66}, Bug, "Abominable Tarantula", 105, []TraitStruct{
		{Visual, false}, {Sound, true}}},
	{[5]uint16{0, 33, 33, 33, 0}, Bug, "Magma Beetle", 124, []TraitStruct{}},
	{[5]uint16{96, 48, 0, 0, 0}, Bug, "Pegasus Mite", 134, []TraitStruct{
		{Sensation, false}, {Sound, false}}},
	{[5]uint16{24, 24, 32, 32, 0}, Bug, "Avalanche Cricket", 140, []TraitStruct{
		{Taste, true}, {Sensation, false}}},
	{[5]uint16{33, 33, 44, 44, 0}, Bug, "Frost Hopper", 196, []TraitStruct{
		{Taste, true}, {Sensation, false}}},

	// Fish
	{[5]uint16{8, 0, 0, 0, 0}, Fish, "River Calamari", 5, []TraitStruct{
		{Sensation, false}}},
	{[5]uint16{24, 0, 0, 0, 0}, Fish, "Swamp Octopus", 18, []TraitStruct{
		{Sensation, false}}},
	{[5]uint16{12, 0, 0, 6, 0}, Fish, "Swamp Fish", 22, []TraitStruct{}},
	{[5]uint16{6, 0, 12, 0, 0}, Fish, "Mud Shrimp", 26, []TraitStruct{
		{Aroma, true}}},
	{[5]uint16{40, 0, 0, 0, 0}, Fish, "Dwarf Kraken", 30, []TraitStruct{
		{Sensation, false}}},
	{[5]uint16{10, 10, 10, 0, 0}, Fish, "Electrocution Eel", 45, []TraitStruct{
		{Visual, true}}},
	{[5]uint16{10, 0, 20, 0, 0}, Fish, "Cobweb Crayfish", 48, []TraitStruct{
		{Aroma, true}}},
	{[5]uint16{0, 0, 0, 0, 32}, Fish, "Crag Crab", 75, []TraitStruct{
		{Aroma, true}}},
	{[5]uint16{24, 24, 24, 0, 0}, Fish, "Hangman Eel", 95, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 0, 44}, Fish, "Blackfrost Lobster", 104, []TraitStruct{
		{Aroma, true}}},
	{[5]uint16{96, 0, 48, 0, 0}, Fish, "Buoyant Blowfish", 138, []TraitStruct{
		{Visual, false}, {Sound, false}}},
	{[5]uint16{132, 0, 66, 0, 0}, Fish, "Icicle Pufferfish", 210, []TraitStruct{
		{Visual, false}, {Sound, false}}},

	// Flesh
	{[5]uint16{0, 8, 0, 0, 0}, Flesh, "Serpent's Slippery Tongue", 6, []TraitStruct{
		{Aroma, false}}},
	{[5]uint16{0, 24, 0, 0, 0}, Flesh, "Salamander's Fiery Tongue", 22, []TraitStruct{
		{Aroma, false}}},
	{[5]uint16{0, 40, 0, 0, 0}, Flesh, "Banshee's Bloody Tongue", 32, []TraitStruct{
		{Aroma, false}}},
	{[5]uint16{0, 0, 24, 12, 0}, Flesh, "Frog Leg", 33, []TraitStruct{
		{Visual, false}}},
	{[5]uint16{0, 16, 0, 0, 0}, Flesh, "Eye of Newt", 34, []TraitStruct{
		{Taste, true}, {Sensation, true}}},
	{[5]uint16{0, 0, 30, 0, 10}, Flesh, "Thunderbird's Molted Feather", 60, []TraitStruct{}},
	{[5]uint16{16, 0, 0, 32, 0}, Flesh, "Harpy's Heart of Stone", 76, []TraitStruct{
		{Sensation, true}}},
	{[5]uint16{0, 0, 0, 48, 16}, Flesh, "Lamia's Shed Scales", 110, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 66, 22}, Flesh, "Body Snatcher's Sloughed Skin", 132, []TraitStruct{}},

	// Bone
	{[5]uint16{0, 0, 8, 0, 0}, Bone, "Unicorn Horn", 6, []TraitStruct{
		{Taste, false}}},
	{[5]uint16{0, 0, 24, 0, 0}, Bone, "Qilin's Tri-Horn", 18, []TraitStruct{
		{Taste, false}}},
	{[5]uint16{6, 0, 12, 0, 0}, Bone, "Crocodile Tooth", 20, []TraitStruct{}},
	{[5]uint16{9, 9, 9, 0, 0}, Bone, "Hydra Vertebra", 35, []TraitStruct{}},
	{[5]uint16{0, 0, 40, 0, 0}, Bone, "Spriggan Antler", 38, []TraitStruct{
		{Taste, false}}},
	{[5]uint16{0, 30, 0, 0, 10}, Bone, "Barghast Canine", 55, []TraitStruct{}},
	{[5]uint16{0, 0, 64, 0, 0}, Bone, "Silver Stag Antler", 72, []TraitStruct{
		{Taste, false}}},
	{[5]uint16{0, 48, 0, 0, 16}, Bone, "Naga's Fang", 98, []TraitStruct{
		{Sensation, false}, {Sound, true}}},
	{[5]uint16{0, 0, 55, 55, 22}, Bone, "Draugr's Tibia", 184, []TraitStruct{}},
	{[5]uint16{0, 0, 40, 40, 16}, Bone, "Stalking Skeleton's Fibula", 150, []TraitStruct{}},

	// Mineral
	{[5]uint16{4, 4, 0, 0, 0}, Mineral, "River-Pixie's Shell", 11, []TraitStruct{}},
	{[5]uint16{12, 12, 0, 0, 0}, Mineral, "Leech Snail's Shell", 26, []TraitStruct{}},
	{[5]uint16{18, 12, 0, 10, 0}, Mineral, "Golemite", 38, []TraitStruct{}},
	{[5]uint16{20, 20, 0, 0, 0}, Mineral, "Slapping Turtle's Shell", 46, []TraitStruct{}},
	{[5]uint16{30, 0, 0, 0, 10}, Mineral, "Sea Salt", 55, []TraitStruct{}},
	{[5]uint16{24, 0, 0, 0, 8}, Mineral, "Rock Salt", 68, []TraitStruct{
		{Taste, true}, {Visual, true}}},
	{[5]uint16{32, 32, 0, 0, 0}, Mineral, "Scimitar Crab's Shell", 76, []TraitStruct{}},
	{[5]uint16{30, 20, 0, 0, 10}, Mineral, "Abyssalite", 79, []TraitStruct{}},
	{[5]uint16{48, 32, 0, 16, 0}, Mineral, "Supernalite", 134, []TraitStruct{
		{Taste, false}, {Visual, true}}},
	{[5]uint16{55, 55, 0, 22, 0}, Mineral, "Hoarite", 167, []TraitStruct{
		{Taste, false}, {Visual, true}}},

	// Essence
	{[5]uint16{4, 0, 4, 0, 0}, Essence, "Kappa Pheromones", 13, []TraitStruct{}},
	{[5]uint16{12, 0, 12, 0, 0}, Essence, "Warg Pheromones", 26, []TraitStruct{}},
	{[5]uint16{20, 0, 20, 0, 0}, Essence, "Nessie Pheromones", 50, []TraitStruct{}},
	{[5]uint16{0, 10, 12, 18, 0}, Essence, "Raven's Shadow", 52, []TraitStruct{}},
	{[5]uint16{0, 0, 30, 10, 0}, Essence, "Raiju Droppings", 55, []TraitStruct{}},
	{[5]uint16{32, 0, 32, 0, 0}, Essence, "Ogre's Shadow", 74, []TraitStruct{}},
	{[5]uint16{0, 0, 30, 20, 10}, Essence, "Dropspider's Shadow", 90, []TraitStruct{}},
	{[5]uint16{0, 0, 64, 32, 0}, Essence, "Chimera Waste", 118, []TraitStruct{
		{Aroma, false}}},
	{[5]uint16{0, 33, 33, 11, 0}, Essence, "Dragon Tear", 118, []TraitStruct{
		{Sensation, true}}},
	{[5]uint16{0, 48, 32, 16, 0}, Essence, "Bioplasm", 125, []TraitStruct{
		{Visual, false}, {Sound, true}}},
	{[5]uint16{0, 55, 55, 22, 0}, Essence, "Xenoplasm", 166, []TraitStruct{
		{Visual, false}, {Sound, true}}},

	// Gem
	{[5]uint16{0, 4, 4, 0, 0}, Gem, "Pixiedust Diamond", 14, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 12, 0}, Gem, "Murkwater Pearl", 27, []TraitStruct{
		{Visual, true}}},
	{[5]uint16{0, 12, 12, 0, 0}, Gem, "Golem's-Eye Diamond", 28, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 20, 0}, Gem, "Shadowveil Pearl", 38, []TraitStruct{
		{Visual, true}}},
	{[5]uint16{0, 20, 20, 0, 0}, Gem, "Spider's-Bait Diamond", 50, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 32, 0}, Gem, "Lustrous Pearl", 60, []TraitStruct{
		{Visual, true}}},
	{[5]uint16{30, 10, 20, 0, 0}, Gem, "Thunder Quartz", 72, []TraitStruct{}},
	{[5]uint16{0, 0, 0, 44, 0}, Gem, "Dragonfire Pearl", 76, []TraitStruct{
		{Visual, true}}},
	{[5]uint16{0, 32, 32, 0, 0}, Gem, "Griffin's-Whetstone Diamond", 86, []TraitStruct{}},
	{[5]uint16{0, 44, 44, 0, 0}, Gem, "Direwolf's-Breath Diamond", 118, []TraitStruct{}},
	{[5]uint16{64, 48, 0, 32, 0}, Gem, "Poison Quartz", 185, []TraitStruct{
		{Sound, false}}},

	// Ore
	{[5]uint16{0, 0, 0, 18, 0}, Ore, "Glass Ore", 24, []TraitStruct{}},
	{[5]uint16{0, 12, 0, 0, 0}, Ore, "Desert Metal", 25, []TraitStruct{
		{Sensation, true}}},
	{[5]uint16{0, 0, 0, 30, 0}, Ore, "Fulgurite Ore", 40, []TraitStruct{}},
	{[5]uint16{0, 0, 16, 0, 0}, Ore, "Celestial Ore", 45, []TraitStruct{
		{Aroma, true}, {Visual, true}}},
	{[5]uint16{0, 0, 22, 0, 0}, Ore, "Nether Ore", 51, []TraitStruct{
		{Aroma, true}, {Visual, true}}},
	{[5]uint16{0, 32, 64, 0, 0}, Ore, "Weeping Metal Ore", 66, []TraitStruct{
		{Sensation, false}}},
	{[5]uint16{30, 10, 0, 0, 20}, Ore, "Malachite Ore", 93, []TraitStruct{}},
	{[5]uint16{66, 66, 0, 0, 33}, Ore, "Amethyst Ore", 206, []TraitStruct{
		{Sound, false}}},

	// Pure Mana
	{[5]uint16{15, 15, 15, 15, 15}, PureMana, "Mote of Mana", 130, []TraitStruct{}},
	{[5]uint16{24, 24, 24, 24, 24}, PureMana, "Ember of Mana", 165, []TraitStruct{}},
	{[5]uint16{50, 40, 30, 20, 10}, PureMana, "Mana Blaze", 234, []TraitStruct{
		{Sensation, true}, {Aroma, true}}},
}

type IngredientCategory byte

const (
	Slime IngredientCategory = iota
	Plant
	Flower
	Fruit
	Fungus
	Bug
	Fish
	Flesh
	Bone
	Mineral
	Essence
	Gem
	Ore
	PureMana
)

type TraitStruct struct {
	Trait  TraitType
	IsGood bool
}

type TraitType byte

func (tt *TraitType) String() string {
	switch *tt {
	case Taste:
		return "Taste"
	case Sensation:
		return "Sensation"
	case Aroma:
		return "Aroma"
	case Visual:
		return "Visual"
	case Sound:
		return "Sound"
	}
	return ""
}

const (
	Taste TraitType = iota
	Sensation
	Aroma
	Visual
	Sound
)

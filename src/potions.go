package src

type Potion struct {
	Magimints [5]uint16
	PK        PotionKind
	Name      string
}

type PotionSearch struct {
	Magimints [5]uint16
}

// todo display potion stars and sizes based on these thresholds
// potions thresholds:
// 	size:		star1		star2		star3		star4		star5		star6
//	minor:	 	0-9, 		10-19, 		20-29, 		30-39, 		40-49,		50-59
//	common: 	60-74, 		75-89, 		90-104, 	105-114,	115-129,	130-149
//	greater:	150-169,	170-194,	195-214,	215-234,	235-259,	260-289
// 	grand:		290-314,	315-344,	345-369,	370-399,	400-429,	430-469
// 	superior:	470-504,	505-544,	545-579,	580-619,	620-659,	660-719
// 	masterwork:	720-

var Potions = []Potion{
	// Potions
	{[5]uint16{1, 1, 0, 0, 0}, PPT, "Health"},
	{[5]uint16{0, 1, 1, 0, 0}, PPT, "Mana"},
	{[5]uint16{1, 0, 0, 0, 1}, PPT, "Stamina"},
	{[5]uint16{0, 0, 1, 1, 0}, PPT, "Speed"},
	{[5]uint16{0, 0, 0, 1, 1}, PPT, "Tolerance"},
	// Tonics
	{[5]uint16{1, 0, 1, 0, 0}, TPT, "Fire"},
	{[5]uint16{1, 0, 0, 1, 0}, TPT, "Ice"},
	{[5]uint16{0, 1, 0, 1, 0}, TPT, "Thunder"},
	{[5]uint16{0, 1, 0, 0, 1}, TPT, "Shadow"},
	{[5]uint16{0, 0, 1, 0, 1}, TPT, "Radiation"},
	// Enhancers
	{[5]uint16{3, 4, 3, 0, 0}, EPT, "Sight"},
	{[5]uint16{0, 3, 4, 3, 0}, EPT, "Alertness"},
	{[5]uint16{4, 3, 0, 0, 3}, EPT, "Insight"},
	{[5]uint16{3, 0, 0, 3, 4}, EPT, "Dowsing"},
	{[5]uint16{0, 0, 3, 4, 3}, EPT, "Seeking"},
	// Cures
	{[5]uint16{2, 0, 1, 1, 0}, CPT, "Poison"},
	{[5]uint16{1, 1, 0, 2, 0}, CPT, "Drowsiness"},
	{[5]uint16{1, 0, 2, 0, 1}, CPT, "Petrification"},
	{[5]uint16{0, 2, 1, 0, 1}, CPT, "Silence"},
	{[5]uint16{0, 1, 1, 0, 2}, CPT, "Curse"},
}

type PotionKind byte

const (
	PPT PotionKind = iota
	TPT
	EPT
	CPT
)

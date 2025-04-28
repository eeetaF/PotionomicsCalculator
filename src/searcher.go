package src

type SearchResult struct {
	ResultingPotion Potion
	Ingredients     map[string]uint16
	TotalMagimints  uint16
	NumberIngreds   uint16
	Traits          []TraitStruct
}

func SearchPerfectCombosByParams(minMags, maxMags, minIngr, maxIngr uint16, desiredPotion string) *[]SearchResult {
	stack := make([]searchUnit, 0, 1000)
	if pt, ok := potionsMapByName[desiredPotion]; ok {
		stack = append(stack, searchUnit{desiredPotion: pt.Magimints})
	} else {
		for mags := range potionsMap {
			stack = append(stack, searchUnit{desiredPotion: mags})
		}
	}
	for len(stack) != 0 {
		currentUnit := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if currentUnit.numIngreds >= maxIngr || currentUnit.numMagimints >= maxMags {
			continue
		}

	}
}

type searchUnit struct {
	desiredPotion [5]uint16
	magimints     [5]uint16
	numIngreds    uint16
	numMagimints  uint16
	ingridients   map[string]uint16
}

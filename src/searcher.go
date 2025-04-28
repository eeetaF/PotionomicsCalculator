package src

type SearchResult struct {
	ResultingPotion Potion
	Ingredients     map[string]uint16
	TotalMagimints  uint16
	NumberIngreds   uint16
	Traits          []TraitStruct // uploads after
}

func SearchPerfectCombosByParams(minMags, maxMags, minIngr, maxIngr uint16, desiredPotion string) *[]SearchResult {
	stack := make([]searchUnit, 0, 1000)
	var i, j uint16
	searchResult := make([]SearchResult, 0, 1000)
	if pt, ok := potionsMapByName[desiredPotion]; ok {
		stack = append(stack, searchUnit{desiredPotion: pt.Magimints, ingredients: make(map[string]uint16, maxIngr)})
	} else {
		for mags := range potionsMap {
			stack = append(stack, searchUnit{desiredPotion: mags, ingredients: make(map[string]uint16, maxIngr)})
		}
	}
	for len(stack) != 0 {
		currentUnit := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if currentUnit.numIngreds > maxIngr || currentUnit.numMagimints > maxMags {
			continue
		}
		magIdx, ok := getNeededMag(&currentUnit.desiredPotion, &currentUnit.magimints)
		if !ok {
			continue
		}
		if magIdx == 5 && currentUnit.numIngreds >= minIngr && currentUnit.numMagimints >= minMags { // good potion
			searchResult = append(searchResult, SearchResult{
				ResultingPotion: potionsMap[currentUnit.desiredPotion],
				Ingredients:     currentUnit.ingredients,
				TotalMagimints:  currentUnit.numMagimints,
				NumberIngreds:   currentUnit.numIngreds,
			})
		}

		numIngredsToAdd := maxIngr - currentUnit.numIngreds
		if numIngredsToAdd == 0 { // no more to add
			continue
		}

		if magIdx == 5 { // mags balanced, we need to add any mag more
			for i = 0; i < 4; i++ {
				if currentUnit.desiredPotion[i] != 0 {
					break
				}
			}
			magIdx = i
		}
		for i = currentUnit.minIngredToUse[magIdx]; int(i) < len(IngredsByMags[magIdx]); i++ {
			ingrInfo := IngredsByMags[magIdx][i]
			ingrName := ingrInfo.name
			if _, ok = currentUnit.ingredients[ingrName]; !ok { // ingred is new, can be added
				for j = 1; j <= numIngredsToAdd; j++ {
					newUnit := searchUnit{
						desiredPotion:  currentUnit.desiredPotion,
						magimints:      currentUnit.magimints,
						numIngreds:     currentUnit.numIngreds,
						numMagimints:   currentUnit.numMagimints,
						minIngredToUse: currentUnit.minIngredToUse,
						ingredients:    make(map[string]uint16, maxIngr),
					}
					newUnit.minIngredToUse[magIdx] = i
					for k, v := range currentUnit.ingredients {
						newUnit.ingredients[k] = v
					}
					addIngred(&ingrInfo, &newUnit, j)
					stack = append(stack, newUnit)
				}
			}
		}
	}
	return &searchResult
}

func addIngred(ingredInfo *NameWithMags, addTo *searchUnit, quantity uint16) {
	for i := range 5 { // add magimints
		(*addTo).magimints[i] += (*ingredInfo).mags[i] * quantity
		(*addTo).numMagimints += (*ingredInfo).mags[i] * quantity
	}
	(*addTo).ingredients[(*ingredInfo).name] += quantity // add ingred to map
	(*addTo).numIngreds += quantity                      // add to total ingreds
}

func getNeededMag(desiredPotion *[5]uint16, magimints *[5]uint16) (uint16, bool) {
	var maxMagIdx, maxMagValue uint16 = 100, 0
	for i, mag := range magimints {
		if mag > 0 && (*desiredPotion)[i] == 0 {
			return 0, false
		}
		if mag > maxMagValue {
			maxMagIdx = uint16(i)
			maxMagValue = mag
		}
	}
	if maxMagIdx == 100 {
		for i, val := range desiredPotion {
			if val != 0 {
				return uint16(i), true
			}
		}
	}
	delim := (*desiredPotion)[maxMagIdx]
	if (*magimints)[maxMagIdx]%delim != 0 {
		return maxMagIdx, true
	}
	unit := (*magimints)[maxMagIdx] / delim
	for i, mag := range *magimints {
		if (*desiredPotion)[i]*unit > mag {
			return uint16(i), true
		} else if (*desiredPotion)[i]*unit < mag {
			return maxMagIdx, true
		}
	}
	return 5, true
}

type searchUnit struct {
	desiredPotion  [5]uint16
	magimints      [5]uint16
	numIngreds     uint16
	numMagimints   uint16
	ingredients    map[string]uint16
	minIngredToUse [5]uint16
}

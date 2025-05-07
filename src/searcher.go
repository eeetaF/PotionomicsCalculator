package src

import (
	"log"
	"sort"
)

const MaxSearchResults = 65535

type SearchResult struct {
	ResultingPotion Potion
	Ingredients     *[]IngredientWithQuantity
	TotalMagimints  uint16
	NumberIngreds   uint16
	Traits          []TraitStruct
}

type IngredientWithQuantity struct {
	Quantity   uint16
	Ingredient Ingredient
}

type SearchOpts struct {
	minMags          uint16
	maxMags          uint16
	minIngr          uint16
	maxIngr          uint16
	topResultsToShow uint16
	desiredPotions   []string
	traits           [5]int8 // -1 for all, 0 excludes bad trait, 1 includes good trait
}

func SearchPerfectCombosByParams(opts *SearchOpts) *[]*SearchResult {
	if len(opts.desiredPotions) == 0 {
		return &[]*SearchResult{}
	}
	opts.minIngr = max(opts.minIngr, 1)

	// these vars are used in calculations. All of them are being constantly reused.
	printThreshold := 1
	var result, lastI, j byte
	var i, newTotalMags, delim, minMagsLocal, maxMagsLocal, numMags, numIngreds, ingredsToAdd, lenLocalMagsSlice uint16
	var isLastI, isLastJ bool
	var newMags [5]uint16
	var localMags *[5]uint16
	var localMagsSlice []*[5]uint16
	var pot potionForSearch
	var localTraits [5]int8
	var currIngTrait int8
	var cu *SearchUnit

	opts.topResultsToShow = min(opts.topResultsToShow, MaxSearchResults)
	results := make([]*SearchResult, 0, opts.topResultsToShow*2)

	neededGoodTraits := completeNeededGoodTraits(&opts.traits) // shared among all potions

	for _, p := range opts.desiredPotions { //  loop - potions
		pot = potionsForSearchMap[p]
		delim = pot.delim
		// make sure that maxMagsLocal is a multiple of delim
		minMagsLocal = max(opts.minMags, delim)
		maxMagsLocal = (opts.maxMags / delim) * delim
		localIngredsByMags := getIngredientDuringSearch(&pot, &opts.traits)
		lastI = getLastI(&pot)

		lenLocalMagsSlice = (maxMagsLocal-minMagsLocal)/delim + 1
		localMagsSlice = make([]*[5]uint16, 0, lenLocalMagsSlice)
		for i = range lenLocalMagsSlice {
			localMagsSlice = append(localMagsSlice, &[5]uint16{})
			for j = range lastI + 1 {
				localMagsSlice[i][j] = pot.mags[j] * (minMagsLocal/delim + i)
			}
		}

		for numIngreds = opts.maxIngr; numIngreds >= opts.minIngr; numIngreds-- { // loop - ingredients
			for numMags = maxMagsLocal; numMags >= minMagsLocal; numMags -= min(delim, numMags) { // loop - magimints

				// here, we finally have fixed target potion, num of mags and num of ingredients.
				localMags = localMagsSlice[(numMags-minMagsLocal)/delim]
				// I can't prove why, but max stack's len has never exceeded numIngreds
				stack := make([]*SearchUnit, 0, numIngreds)
				stack = append(stack, &SearchUnit{ingredsUsed: make([]usedIngredient, 0, opts.maxIngr)})

				for len(stack) > 0 {
					cu = stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					isLastJ = true

					for cu.i != lastI+1 && uint16(len(localIngredsByMags[cu.i])) == cu.j { // skip finished mags
						cu.i++
						cu.j = 0
						isLastJ = false
					}
					if cu.i > lastI { // skip last
						continue
					}
					if isLastJ {
						isLastJ = uint16(len(localIngredsByMags[cu.i])) == cu.j+1
					}
					isLastI = cu.i == lastI

					if cu.mags[cu.i]+localIngredsByMags[cu.i][cu.j].mags[cu.i] > localMags[cu.i] {
						continue // current ingredient exceeds the number of mags remained, no more such ingreds to add
					}

					if !isLastJ { // there are more ingreds with the same mag type, we may skip this ingr
						newCu := &SearchUnit{
							i:            cu.i,
							j:            cu.j + 1,
							mags:         cu.mags, // safe, copies the array
							ingredsUsed:  make([]usedIngredient, len(cu.ingredsUsed), cap(cu.ingredsUsed)),
							totalMags:    cu.totalMags,
							totalIngreds: cu.totalIngreds,
						}
						copy(newCu.ingredsUsed, cu.ingredsUsed)
						stack = append(stack, newCu)
					}

					// at this point, i and j are valid according to localIngredsByMags
					for ingredsToAdd = 1; ingredsToAdd <= (numIngreds - cu.totalIngreds); ingredsToAdd++ {
						if localIngredsByMags[cu.i][cu.j].quantityAvailable < ingredsToAdd {
							break // not enough such ingreds left
						}

						newTotalMags = cu.totalMags + ingredsToAdd*localIngredsByMags[cu.i][cu.j].totalMags
						if newTotalMags > numMags || // total mags is bigger than numMags
							(newTotalMags < numMags && ingredsToAdd == (numIngreds-cu.totalIngreds)) { // total mags is less and no more ingreds to add
							break
						}
						if newTotalMags != numMags && isLastI && isLastJ { // mags not equal, but ingred is the last in both lists
							continue
						}

						newMags = [5]uint16{}
						for j = range lastI + 1 {
							newMags[j] = cu.mags[j] + ingredsToAdd*localIngredsByMags[cu.i][cu.j].mags[j]
						}

						// 0 - default, 1 - bad, 2 - mag finished, 3 - good. IMPORTANT: regards only mags. Not traits / numIngreds etc.
						result = getNewMagsStatus(&newMags, localMags, cu.i, lastI)

						if result == 1 {
							break
						}
						if result == 0 && isLastJ { // last elem in the list and current mag is not finished, need to add more such ingred
							continue
						}
						if result == 3 {
							// need to check for good traits, numIngreds
							if ingredsToAdd+cu.totalIngreds != numIngreds {
								break
							}
							// setup resulting traits
							localTraits = getTraits(&cu.ingredsUsed, localIngredsByMags)
							for traitIdx, trait := range localTraits {
								if trait != -1 {
									currIngTrait = localIngredsByMags[cu.i][cu.j].traits[traitIdx]
									if currIngTrait == -1 {
										localTraits[traitIdx] = -1
										continue
									}
									localTraits[traitIdx] = max(trait, currIngTrait)
								}
							}
							// check if all resulting traits are valid according to the search
							for _, trait := range *neededGoodTraits {
								if localTraits[trait] != 1 {
									result = 1
									break
								}
							}
							if result == 1 {
								break // bad traits
							}

							// all good, add to results
							cu.ingredsUsed = append(cu.ingredsUsed, usedIngredient{
								i:        cu.i,
								j:        cu.j,
								quantity: ingredsToAdd,
							})

							trts := make([]TraitStruct, 0, 5)
							for trtIdx, val := range localTraits {
								if val == -1 {
									trts = append(trts, TraitStruct{TraitType(trtIdx), false})
								} else if val == 1 {
									trts = append(trts, TraitStruct{TraitType(trtIdx), true})
								}
							}
							results = append(results, &SearchResult{
								ResultingPotion: potionsMap[p],
								Ingredients:     usedIngredientsToSortedIngredients(&cu.ingredsUsed, localIngredsByMags),
								TotalMagimints:  numMags,
								NumberIngreds:   numIngreds,
								Traits:          trts,
							})
							if len(results)%printThreshold == 0 {
								if len(results) == printThreshold*10 {
									printThreshold *= 10
								}
								log.Printf("[INFO] Recipes found: %d\n", len(results))
								if len(results) > int(10*opts.topResultsToShow) {
									log.Printf("[INFO] found 10x similar recipes, leaving serch")
									log.Printf("[RSLT] Total recipes found: %d", len(results))
									return &results // enough results found
								}
							}
							break
						}

						// Here, all checks are passed. We just add the ingredient
						newCu := &SearchUnit{
							i:            cu.i,
							j:            cu.j + 1,
							mags:         newMags, // safe, copies the array
							ingredsUsed:  make([]usedIngredient, len(cu.ingredsUsed), cap(cu.ingredsUsed)),
							totalMags:    newTotalMags,
							totalIngreds: cu.totalIngreds + ingredsToAdd,
						}
						if result == 2 {
							newCu.i++
							newCu.j = 0
						}
						copy(newCu.ingredsUsed, cu.ingredsUsed)
						newCu.ingredsUsed = append(newCu.ingredsUsed, usedIngredient{
							cu.i,
							cu.j,
							ingredsToAdd,
						})
						stack = append(stack, newCu)
					}
				}

				if len(results) >= int(opts.topResultsToShow) {
					log.Printf("[INFO] found enough best recipes, leaving search")
					log.Printf("[RSLT] Total recipes found: %d", len(results))
					return &results // enough recipes found
				}
			}
		}
	}
	log.Printf("[INFO] all combinations checked")
	log.Printf("[RSLT] Total recipes found: %d", len(results))
	return &results
}

func usedIngredientsToSortedIngredients(ings *[]usedIngredient, localIngredsByMags *[5][]ingredientDuringSearch) *[]IngredientWithQuantity {
	newIngs := make([]IngredientWithQuantity, 0, len(*ings))

	sort.Slice(*ings, func(i, j int) bool {
		if localIngredsByMags[(*ings)[i].i][(*ings)[i].j].totalMags == localIngredsByMags[(*ings)[j].i][(*ings)[j].j].totalMags {
			// total mags are equal
			return (*ings)[i].quantity > (*ings)[j].quantity
		}
		return localIngredsByMags[(*ings)[i].i][(*ings)[i].j].totalMags > localIngredsByMags[(*ings)[j].i][(*ings)[j].j].totalMags
	})

	for _, ingUsed := range *ings {
		newIngs = append(newIngs, IngredientWithQuantity{
			Quantity:   ingUsed.quantity,
			Ingredient: Ingredients[localIngredsByMags[ingUsed.i][ingUsed.j].id],
		})
	}
	return &newIngs
}

func getTraits(ingredsUsed *[]usedIngredient, localIngredsByMags *[5][]ingredientDuringSearch) [5]int8 {
	var result [5]int8
	for _, ingred := range *ingredsUsed {
		for i, trait := range localIngredsByMags[ingred.i][ingred.j].traits {
			if trait == -1 {
				result[i] = -1
			} else if trait == 1 && result[i] != -1 {
				result[i] = 1
			}
		}
	}
	return result
}

// 0 - default result
// 1 - status bad: any of mags bigger than expected
// 2 - status finished: current magimint finished, but any other is not
// 3 - status good: a solution found, add it to the list
func getNewMagsStatus(newMags *[5]uint16, expectedMags *[5]uint16, currentMag, lastI byte) byte {
	good, finished := true, false

	if newMags[currentMag] > expectedMags[currentMag] {
		return 1
	}
	if newMags[currentMag] == expectedMags[currentMag] {
		finished = true
	}
	currentMag++

	for ; currentMag <= lastI; currentMag++ {
		if newMags[currentMag] > expectedMags[currentMag] {
			return 1
		}
		if newMags[currentMag] < expectedMags[currentMag] {
			good = false
		}
	}
	if !finished {
		return 0
	}
	if good {
		return 3
	}
	return 2
}

func getLastI(p *potionForSearch) byte {
	var i byte
	for i = 4; i != 0; i-- {
		if p.mags[i] != 0 {
			return i
		}
	}
	return 0
}

func completeNeededGoodTraits(traits *[5]int8) *[]TraitType {
	var neededGoodTraits []TraitType
	for i := range traits {
		if traits[i] == 1 {
			neededGoodTraits = append(neededGoodTraits, TraitType(i))
		}
	}
	return &neededGoodTraits
}

func getIngredientDuringSearch(p *potionForSearch, traits *[5]int8) *[5][]ingredientDuringSearch {
	var ingredientsByMags [5][]ingredientDuringSearch
	for i := range ingredientsByLimitedMagsSetup {
		if p.mags[i] == 0 {
			continue
		}
		for _, ingred := range ingredientsByLimitedMagsSetup[i] {
			var badIngred bool
			for j := i; j < 5; j++ {
				if ingred.mags[j] != 0 && p.mags[j] == 0 {
					badIngred = true
					break
				}
			}
			if !badIngred {
				for j := range ingred.traits {
					if (*traits)[j] >= 0 && ingred.traits[j] == -1 {
						badIngred = true
						break
					}
				}
			}
			if badIngred {
				continue
			}
			// ingred is good, can be added
			ingredientsByMags[i] = append(ingredientsByMags[i], ingredientDuringSearch{
				id:                ingred.id,
				traits:            ingred.traits,
				mags:              ingred.mags,
				totalMags:         ingred.totalMags,
				quantityAvailable: ingred.quantityAvailable,
			})
		}
	}
	return &ingredientsByMags
}

type SearchUnit struct {
	i            byte   // indexes in localIngredsByMags
	j            uint16 // indexes in localIngredsByMags
	mags         [5]uint16
	ingredsUsed  []usedIngredient
	totalMags    uint16
	totalIngreds uint16
}

type usedIngredient struct {
	i        byte   // first index in ingredientDuringSearch
	j        uint16 // second index in ingredientDuringSearch
	quantity uint16
}

type ingredientDuringSearch struct {
	id                uint16    // uniquely identifies the ingredient in Ingredients
	traits            [5]int8   // -1 for bad traits, 1 for good traits, 0 for no traits
	mags              [5]uint16 // mags of ingredient
	totalMags         uint16
	quantityAvailable uint16
}

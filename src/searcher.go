package src

import (
	"context"
	"fmt"
	"runtime"
	"sort"
	"sync"
	"time"
)

const MaxSearchResults = 200

type SearchResult struct {
	ResultingPotion Potion
	Ingredients     []nameWithQuantity
	TotalMagimints  uint16
	NumberIngreds   uint16
	Traits          []TraitStruct // uploads after
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

func SearchPerfectCombosByParams(opts *SearchOpts) *[]SearchResult {
	if len(opts.desiredPotions) == 0 {
		return &[]SearchResult{}
	}

	results := make([]SearchResult, 0, MaxSearchResults)

	neededGoodTraits := completeNeededGoodTraits(&opts.traits) // shared among all potions

	for _, p := range opts.desiredPotions { //  loop - potions
		delim := potionsForSearchMap[p].delim
		// make sure that maxMagsLocal is a multiple of delim
		maxMagsLocal := opts.maxMags / delim * delim
		localIngredsByMags := getIngredientDuringSearch(potionsForSearchMap[p], &opts.traits)

		for numIngreds := opts.maxIngr; numIngreds >= opts.minIngr && numIngreds != 65535; numIngreds-- { // loop - ingredients
			for numMags := maxMagsLocal; numMags >= opts.minMags; numMags -= min(delim, numMags) { // loop - magimints

				// here, we finally have fixed target potion, num of mags and num of ingredients.
				stack := make([]SearchUnit, 0, 1000)
				stack = append(stack, completeSearchUnit())

				for len(stack) > 0 {
					cu := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

				}
			}
		}
	}
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

func completeSearchUnit() SearchUnit {
	var quantityAvailable [5][]uint16
	for i := range ingredientsByLimitedMagsSetup {
		for _, ingred := range ingredientsByLimitedMagsSetup[i] {
			quantityAvailable[i] = append(quantityAvailable[i], ingred.quantityAvailable)
		}
	}

	return SearchUnit{quantityAvailable: quantityAvailable}
}

func getIngredientDuringSearch(p potionForSearch, traits *[5]int8) *[5][]ingredientDuringSearch {
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
					if (*traits)[j] == 1 && ingred.traits[j] == -1 {
						badIngred = true
						break
					}
				}
			}
			if badIngred {
				continue
			}
			// ingred is good, can be added
			ingDuringSearch := ingredientDuringSearch{
				id:        ingred.id,
				traits:    ingred.traits,
				mags:      ingred.mags,
				totalMags: ingred.totalMags,
			}
			ingredientsByMags[i] = append(ingredientsByMags[i], ingDuringSearch)
		}
	}
	return &ingredientsByMags
}

type SearchUnit struct {
	i, j              uint16 // indexes in localIngredsByMags
	quantityAvailable [5][]uint16
	ingredsUsed       []usedIngredient
	totalMags         uint16
	totalIngreds      uint16
}

type usedIngredient struct {
	i        uint16 // first index in ingredientDuringSearch
	j        uint16 // second index in ingredientDuringSearch
	quantity uint16
}

type ingredientDuringSearch struct {
	id        uint16    // uniquely identifies the ingredient in Ingredients
	traits    [5]int8   // -1 for bad traits, 1 for good traits, 0 for no traits
	mags      [5]uint16 // mags of ingredient
	totalMags uint16
}

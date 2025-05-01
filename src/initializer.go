package src

import (
	"bufio"
	"fmt"
)

var ingredientsMap map[uint16]Ingredient
var potionsForSearchMap map[string]potionForSearch

// 0: all ingreds that have A magimint in it
// 1: all ingreds that don't have A magimint in it
// 2: all ingreds that don't have A, B magimint in it
// 3: all ingreds that don't have A, B, C magimint in it
// 4: all ingreds that don't have A, B, C, D magimint in it
var ingredientsByLimitedMagsSetup [5][]ingredientBeforeSearchSetup

type ingredientBeforeSearchSetup struct {
	id                uint16    // uniquely identifies the ingredient
	traits            [5]int8   // -1 for bad traits, 1 for good traits, 0 for no traits
	mags              [5]uint16 // mags of ingredient
	totalMags         uint16
	quantityAvailable uint16 // number of ingreds available
}

type potionForSearch struct {
	delim uint16 // equals to a sum of magimint units needed for potion
	mags  [5]uint16
}

var In *bufio.Reader
var Out *bufio.Writer

func Initialize() {
	ingredientsMap = make(map[uint16]Ingredient, len(Ingredients))
	for ingIndex, ing := range Ingredients {
		// fill in ingredientsMap
		ingredientsMap[uint16(ingIndex)] = ing

		// fill in ingredientsByLimitedMagsSetup
		var indexToPut byte
		var traits [5]int8
		var totalMags uint16
		for _, mag := range ing.Magimints {
			if mag != 0 {
				totalMags += mag
				break
			}
			indexToPut++
		}
		for i := indexToPut + 1; i < 5; i++ {
			totalMags += ing.Magimints[i]
		}
		for _, trait := range ing.Traits {
			if trait.IsGood {
				traits[trait.Trait] = 1
			} else {
				traits[trait.Trait] = -1
			}
		}
		ingredientsByLimitedMagsSetup[indexToPut] = append(ingredientsByLimitedMagsSetup[indexToPut], ingredientBeforeSearchSetup{
			id:                uint16(ingIndex),
			traits:            traits,
			mags:              ing.Magimints,
			totalMags:         totalMags,
			quantityAvailable: 65535, // for now, keep limitless
		})
	}
	for _, potion := range Potions {
		potionsForSearchMap[potion.Name] = potionForSearch{
			delim: potion.Magimints[0] + potion.Magimints[1] + potion.Magimints[2] + potion.Magimints[3] + potion.Magimints[4],
			mags:  potion.Magimints,
		}
	}
}

func PrintWithBufio(s string) {
	fmt.Fprint(Out, s)
	Out.Flush()
}

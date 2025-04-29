package src

import (
	"bufio"
	"fmt"
)

var ingredientsMap map[string]Ingredient
var potionsMap map[[5]uint16]Potion
var potionsMapByName map[string]Potion

var IngredsByMags [5][]NameWithMagsAndNegTraits

type NameWithMagsAndNegTraits struct {
	name           string
	mags           [5]uint16
	negativeTraits *[5]bool
}

type NameWithMags struct {
	name string
	mags [5]uint16
}

var In *bufio.Reader
var Out *bufio.Writer

func Initialize() {
	ingredientsMap = make(map[string]Ingredient, len(Ingredients))
	potionsMap = make(map[[5]uint16]Potion, len(Potions))
	potionsMapByName = make(map[string]Potion, len(Potions))
	for i := range 5 {
		IngredsByMags[i] = make([]NameWithMagsAndNegTraits, 0, len(Ingredients))
	}

	for _, ing := range Ingredients {
		var negativeTraits [5]bool
		var needToAddTrait bool
		for _, trt := range ing.Traits {
			if !trt.IsGood {
				negativeTraits[trt.Trait] = true
				needToAddTrait = true
			}
		}
		ingredientsMap[ing.Name] = ing
		for i, mag := range ing.Magimints {
			if mag != 0 {
				if needToAddTrait {
					IngredsByMags[i] = append(IngredsByMags[i], NameWithMagsAndNegTraits{ing.Name, ing.Magimints, &negativeTraits})
				} else {
					IngredsByMags[i] = append(IngredsByMags[i], NameWithMagsAndNegTraits{ing.Name, ing.Magimints, nil})
				}
			}
		}
	}

	for _, pot := range Potions {
		potionsMap[pot.Magimints] = pot
		potionsMapByName[pot.Name] = pot
	}
}

func PrintWithBufio(s string) {
	fmt.Fprint(Out, s)
	Out.Flush()
}

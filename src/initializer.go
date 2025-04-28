package src

import (
	"bufio"
	"fmt"
)

var ingredientsMap map[string]Ingredient
var potionsMap map[[5]uint16]Potion
var potionsMapByName map[string]Potion

var PotionsByMags [5][]NameWithMags

type NameWithMags struct {
	name string
	mags [5]uint16
}

var In *bufio.Reader
var Out *bufio.Writer

func Initialize() {
	ingredientsMap = make(map[string]Ingredient, len(Ingredients))
	potionsMap = make(map[[5]uint16]Potion, len(Potions))
	for i := range 5 {
		PotionsByMags[i] = make([]NameWithMags, 0, len(Potions))
	}

	for _, ing := range Ingredients {
		ingredientsMap[ing.Name] = ing
		for i, mag := range ing.Magimints {
			if mag != 0 {
				PotionsByMags[i] = append(PotionsByMags[i], NameWithMags{ing.Name, ing.Magimints})
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

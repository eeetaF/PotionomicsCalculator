package src

import (
	"bufio"
	"fmt"
)

var ingredientsMap map[string]Ingredient
var potionsMap map[[5]uint16]Potion

var In *bufio.Reader
var Out *bufio.Writer

func Initialize() {
	ingredientsMap = make(map[string]Ingredient, len(Ingredients))
	potionsMap = make(map[[5]uint16]Potion, len(Potions))

	for _, ing := range Ingredients {
		ingredientsMap[ing.Name] = ing
	}

	for _, pot := range Potions {
		potionsMap[pot.Magimints] = pot
	}
}

func PrintWithBufio(s string) {
	fmt.Fprint(Out, s)
	Out.Flush()
}

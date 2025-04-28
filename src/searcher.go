package src

type SearchResult struct {
	ResultingPotion Potion
	Ingredients     []Ingredient
	TotalMagimints  uint16
	NumberIngreds   uint16
	Traits          []TraitStruct
}

func SearchPerfectCombosByParams(minMags, maxMags, minIngr, maxIngr uint16, desiredPotion string) *[]SearchResult {

}

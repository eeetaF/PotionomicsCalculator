package src

import (
	"fmt"
	"sort"
)

func SortSearchResult(sr *[]SearchResult) {
	sort.Slice(*sr, func(i, j int) bool {
		return (*sr)[i].TotalMagimints > (*sr)[j].TotalMagimints
	})
	for _, searchResult := range *sr {
		sort.Slice(searchResult.Ingredients, func(i, j int) bool {
			return searchResult.Ingredients[i].Price > searchResult.Ingredients[j].Price
		})
	}
}
func PrintSearchResult(sr *[]SearchResult) {
	var s string
	for _, searchResult := range *sr {
		s += fmt.Sprintf("%d %d %s\n\t", searchResult.NumberIngreds, searchResult.TotalMagimints, searchResult.ResultingPotion.Name)
		for _, trait := range searchResult.Traits {
			if trait.IsGood {
				s += fmt.Sprintf("+%s ", trait.Trait)
			}
		}
		if len(searchResult.Traits) != 0 {
			s += "\n\t"
		}
	}

}

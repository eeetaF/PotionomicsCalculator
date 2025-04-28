package src

import (
	"fmt"
	"sort"
)

func SortSearchResult(sr *[]SearchResult, topResultsToShow uint16) {
	sort.Slice(*sr, func(i, j int) bool {
		return (*sr)[i].TotalMagimints > (*sr)[j].TotalMagimints
	})
	*sr = (*sr)[:min(int(topResultsToShow), len(*sr))]
}
func PrintSearchResult(sr *[]SearchResult) {
	var s string
	for _, searchResult := range *sr {
		s += fmt.Sprintf("%s, ingredients: %d, magimints: %d", searchResult.ResultingPotion.Name, searchResult.NumberIngreds, searchResult.TotalMagimints)
		if len(searchResult.Traits) != 0 {
			s += "\n\t"
		}
		for _, trait := range searchResult.Traits {
			if trait.IsGood {
				s += fmt.Sprintf("+%s ", trait.Trait.String())
			} else {
				s += fmt.Sprintf("-%s ", trait.Trait.String())
			}
		}
		for ingName, ingCount := range searchResult.Ingredients {
			mags := ingredientsMap[ingName].Magimints
			s += fmt.Sprintf("\n\tx%d: [%d, %d, %d, %d, %d] %s",
				ingCount, mags[0], mags[1], mags[2], mags[3], mags[4], ingName)
		}
		s += "\n"
	}
	s += "\n"
	PrintWithBufio(s)
}

package src

import (
	"fmt"
	"sort"
)

func SortSearchResult(sr *[]SearchResult) {
	sort.Slice(*sr, func(i, j int) bool {
		return (*sr)[i].TotalMagimints > (*sr)[j].TotalMagimints
	})
}
func PrintSearchResult(sr *[]SearchResult) {
	var s string
	for _, searchResult := range *sr {
		s += fmt.Sprintf("%d %d %s", searchResult.NumberIngreds, searchResult.TotalMagimints, searchResult.ResultingPotion.Name)
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
			s += fmt.Sprintf("\n\tx%d: %s", ingCount, ingName)
		}
		s += "\n"
	}
	PrintWithBufio(s)
}

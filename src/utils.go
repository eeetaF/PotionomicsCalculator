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
	CompleteTraits(sr)
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

func CompleteTraits(sr *[]SearchResult) {
	var traits [5]byte
	for i := range *sr {
		traits = [5]byte{}
		for ingred := range (*sr)[i].Ingredients {
			ingredInfo := ingredientsMap[ingred]
			for _, traitInfo := range ingredInfo.Traits {
				if !traitInfo.IsGood {
					traits[traitInfo.Trait] = 2
				} else if traitInfo.IsGood && traits[traitInfo.Trait] == 0 {
					traits[traitInfo.Trait] = 1
				}
			}
		}
		for j, val := range traits {
			if val == 1 {
				(*sr)[i].Traits = append((*sr)[i].Traits, TraitStruct{Trait: TraitType(j), IsGood: true})
			} else if val == 2 {
				(*sr)[i].Traits = append((*sr)[i].Traits, TraitStruct{Trait: TraitType(j), IsGood: false})
			}
		}
	}
}

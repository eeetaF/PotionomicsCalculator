package src

import (
	"fmt"
	"sort"
)

func SortAndFilterSearchResult(sr *[]SearchResult, topResultsToShow uint16) {
	CompleteTraits(sr) // fill the traits of results

	sort.Slice(*sr, func(i, j int) bool {
		if (*sr)[i].ResultingPotion.Name == (*sr)[j].ResultingPotion.Name { // first, try sorting by potion name
			if (*sr)[i].NumberIngreds == (*sr)[j].NumberIngreds { // then, try sorting by number ingreds
				if (*sr)[i].TotalMagimints == (*sr)[j].TotalMagimints { // then, try sorting by number magimints
					numPointsI := getNumPoints(&(*sr)[i])
					numPointsI -= getNumPoints(&(*sr)[j])
					if numPointsI == 0 { // finally, try sorting by total ingreds price
						totalIngredsPrice := getTotalIngredsPrice(&(*sr)[i])
						totalIngredsPrice -= getTotalIngredsPrice(&(*sr)[j])
						return totalIngredsPrice < 0
					}
					return numPointsI > 0
				}
				return (*sr)[i].TotalMagimints > (*sr)[j].TotalMagimints
			}
			return (*sr)[i].NumberIngreds > (*sr)[j].NumberIngreds
		}
		return (*sr)[i].ResultingPotion.Name < (*sr)[j].ResultingPotion.Name
	})
	*sr = (*sr)[:min(int(topResultsToShow), len(*sr))]
}

func getTotalIngredsPrice(sr *SearchResult) int {
	var cost int
	for _, ingr := range sr.Ingredients {
		cost += ingredientsMap[ingr.Name].Price
	}
	return cost
}

func getNumPoints(sr *SearchResult) int8 {
	var numPoints int8
	for _, trait := range sr.Traits {
		if trait.IsGood {
			numPoints++
		} else {
			numPoints--
		}
	}
	return numPoints
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
		for _, ingr := range searchResult.Ingredients {
			mags := ingredientsMap[ingr.Name].Magimints
			s += fmt.Sprintf("\n\tx%d: [%d, %d, %d, %d, %d] %s",
				ingr.Quantity, mags[0], mags[1], mags[2], mags[3], mags[4], ingr.Name)
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
		for _, ingr := range (*sr)[i].Ingredients {
			ingredInfo := ingredientsMap[ingr.Name]
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

func CompleteLocalIndreds(localIngredientsMap *[5][]NameWithMags, traits *[5]bool) {
	for i := range 5 {
		for _, val := range IngredsByMags[i] {
			if val.negativeTraits != nil {
				for j := range 5 { // check all traits
					if val.negativeTraits[j] && traits[j] {
						goto nextIngred
					}
				}
			}
			localIngredientsMap[i] = append(localIngredientsMap[i], NameWithMags{val.name, val.mags})
		nextIngred:
		}
	}
}

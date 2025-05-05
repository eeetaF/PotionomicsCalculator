package src

import (
	"sort"
)

func SortAndFilterSearchResult(sr *[]*SearchResult, topResultsToShow uint16) {
	sort.Slice(*sr, func(i, j int) bool {
		if (*sr)[i].ResultingPotion.Name == (*sr)[j].ResultingPotion.Name { // first, try sorting by potion name
			if (*sr)[i].NumberIngreds == (*sr)[j].NumberIngreds { // then, try sorting by number ingreds
				traitsPoints := getTraitsPoints((*sr)[i]) - getTraitsPoints((*sr)[j])
				if traitsPoints == 0 { // then, try sorting by traits
					if (*sr)[i].TotalMagimints == (*sr)[j].TotalMagimints { // then, try sorting by number magimints
						numPointsI := getNumPoints((*sr)[i])
						numPointsI -= getNumPoints((*sr)[j])
						if numPointsI == 0 { // finally, try sorting by total ingreds price
							totalIngredsPrice := getTotalIngredsPrice((*sr)[i])
							totalIngredsPrice -= getTotalIngredsPrice((*sr)[j])
							return totalIngredsPrice < 0
						}
						return numPointsI > 0
					}
					return (*sr)[i].TotalMagimints > (*sr)[j].TotalMagimints
				}
				return traitsPoints > 0
			}
			return (*sr)[i].NumberIngreds > (*sr)[j].NumberIngreds
		}
		return (*sr)[i].ResultingPotion.Name < (*sr)[j].ResultingPotion.Name
	})
	*sr = (*sr)[:min(int(topResultsToShow), len(*sr))]
}

func getTraitsPoints(sr *SearchResult) int8 {
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

func getTotalIngredsPrice(sr *SearchResult) int {
	var cost int
	for _, ingr := range *sr.Ingredients {
		cost += ingr.Ingredient.Price
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

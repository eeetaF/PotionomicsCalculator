package src

import (
	"fmt"
)

func MainLoop() {
	var minMags, maxMags, minIngr, maxIngr, topResultsToShow, taste, sensation, aroma, visual, sound uint16
	var desiredPotion string
	for {
		PrintWithBufio("Enter filter:\n{minMags, maxMags, minIngr, maxIngr, desiredPotion(empty for all), topResultsToShow, Taste, Sensation, Aroma, Visual, Sound}\n")
		fmt.Fscan(In, &minMags, &maxMags, &minIngr, &maxIngr, &desiredPotion, &topResultsToShow, &taste, &sensation, &aroma, &visual, &sound)
		PrintWithBufio("\n")
		neededTraits := make([]TraitType, 0, taste+sensation+aroma+visual+sound)
		if taste == 1 {
			neededTraits = append(neededTraits, Taste)
		}
		if sensation == 1 {
			neededTraits = append(neededTraits, Sensation)
		}
		if aroma == 1 {
			neededTraits = append(neededTraits, Aroma)
		}
		if visual == 1 {
			neededTraits = append(neededTraits, Visual)
		}
		if sound == 1 {
			neededTraits = append(neededTraits, Sound)
		}

		searchResult := SearchPerfectCombosByParams(minMags, maxMags, minIngr, maxIngr, desiredPotion, &neededTraits)
		SortAndFilterSearchResult(searchResult, topResultsToShow)
		PrintSearchResult(searchResult)
		run(searchResult)

		// TODO add filter by traits (include only positive)
		// TODO secondary sort by traits
		// FIXME fix bug: when encounter a multi-magimints ingredient, the result may double. F.e.:
		// 420 420 8 8 empty 100 1 1 1 1 1 -> results in 2 identical thunder potions
	}
}

package src

import "fmt"

func MainLoop() {
	var minMags, maxMags, minIngr, maxIngr, topResultsToShow uint16
	var desiredPotion string
	for {
		PrintWithBufio("Enter filter: {minMags, maxMags, minIngr, maxIngr, desiredPotion(empty for all), topResultsToShow}\n")
		fmt.Fscan(In, &minMags, &maxMags, &minIngr, &maxIngr, &desiredPotion, &topResultsToShow)
		PrintWithBufio("\n")

		searchResult := SearchPerfectCombosByParams(minMags, maxMags, minIngr, maxIngr, desiredPotion)
		SortSearchResult(searchResult, topResultsToShow)
		PrintSearchResult(searchResult)
	}
}

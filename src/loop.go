package src

import "fmt"

func MainLoop() {
	var minMags, maxMags, minIngr, maxIngr uint16
	var desiredPotion string
	for {
		PrintWithBufio("Enter filter: {minMags, maxMags, minIngr, maxIngr, desiredPotion(empty for all)}\n")
		fmt.Fscan(In, &minMags, &maxMags, &minIngr, &maxIngr, &desiredPotion)
		PrintWithBufio("\n")

		searchResult := SearchPerfectCombosByParams(minMags, maxMags, minIngr, maxIngr, desiredPotion)
		SortSearchResult(searchResult)
		PrintSearchResult(searchResult)
	}
}

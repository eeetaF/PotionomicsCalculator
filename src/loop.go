package src

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func MainLoop() {
	var minMags, maxMags, minIngr, maxIngr, topResultsToShow int
	var taste, sensation, aroma, visual, sound int8
	var desiredPotions string
	for {
		PrintWithBufio("Enter {minMags, maxMags}: ")
		fmt.Fscan(In, &minMags, &maxMags)
		PrintWithBufio("Enter {minIngr, maxIngr}: ")
		fmt.Fscan(In, &minIngr, &maxIngr)
		PrintWithBufio(fmt.Sprintf("Enter {topResultsToShow} (if bigger than %d, will be reduced to it): ", MaxSearchResults))
		fmt.Fscan(In, &topResultsToShow)
		PrintWithBufio("Enter desired potions (type empty to include all). Format: {p1_p2_p3}: ")
		fmt.Fscan(In, &desiredPotions)
		PrintWithBufio("Enter traits (Taste, Sensation, Aroma, Visual, Sound) 1 for good only, 0 excludes bad, -1 for all: ")
		fmt.Fscan(In, &taste, &sensation, &aroma, &visual, &sound)
		PrintWithBufio("----------------------\n")
		log.Println("Starting search...")
		desiredPotionsSlc := strings.Split(desiredPotions, "_")
		goodPotions := make([]string, 0, len(desiredPotionsSlc))
		for _, p := range desiredPotionsSlc {
			if _, ok := potionsMap[p]; ok {
				goodPotions = append(goodPotions, p)
			}
		}
		if len(goodPotions) == 0 {
			for _, p := range Potions {
				goodPotions = append(goodPotions, p.Name)
			}
		}

		start := time.Now()

		searchResult := SearchPerfectCombosByParams(&SearchOpts{
			minMags:          uint16(min(minMags, 65535)),
			maxMags:          uint16(min(maxMags, 65535)),
			minIngr:          uint16(min(minIngr, 65535)),
			maxIngr:          uint16(min(maxIngr, 65535)),
			topResultsToShow: uint16(min(topResultsToShow, 65535)),
			desiredPotions:   goodPotions,
			traits:           [5]int8{taste, sensation, aroma, visual, sound},
		})

		SortAndFilterSearchResult(searchResult, uint16(min(topResultsToShow, 65535)))

		PrintWithBufio(fmt.Sprintf("\nSearch took: %s\n", time.Since(start).String()))
		PrintWithBufio("----------------------\n")

		run(searchResult)
		// todo improve searcher: don't fix numIngreds and numMags
	}
}

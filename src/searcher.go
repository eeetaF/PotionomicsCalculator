package src

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type SearchResult struct {
	ResultingPotion Potion
	Ingredients     map[string]uint16
	TotalMagimints  uint16
	NumberIngreds   uint16
	Traits          []TraitStruct // uploads after
}

var stack []searchUnit
var stackMu sync.Mutex

func SearchPerfectCombosByParams(minMags, maxMags, minIngr, maxIngr uint16, desiredPotion string) *[]SearchResult {
	stack = make([]searchUnit, 0, 10000)
	searchResult := make([]SearchResult, 0, 1000)
	var wg sync.WaitGroup

	start := time.Now()

	// Инициализация начального стека
	if pt, ok := potionsMapByName[desiredPotion]; ok {
		wg.Add(1)
		stack = append(stack, searchUnit{desiredPotion: pt.Magimints})
	} else {
		for mags := range potionsMap {
			wg.Add(1)
			stack = append(stack, searchUnit{desiredPotion: mags})
		}
	}

	var mu sync.Mutex

	numWorkers := runtime.NumCPU()

	ctx, cancelFn := context.WithCancel(context.Background())
	for w := 0; w < numWorkers; w++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					stackMu.Lock()
					if len(stack) == 0 {
						stackMu.Unlock()
						continue
					}
					currentUnit := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					stackMu.Unlock()
					processUnit(currentUnit, minMags, maxMags, minIngr, maxIngr, &searchResult, &mu, &wg)
				}
			}
		}()
	}

	wg.Wait()
	cancelFn()
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	return &searchResult
}

func processUnit(currentUnit searchUnit, minMags, maxMags, minIngr, maxIngr uint16, searchResult *[]SearchResult, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	var i, j uint16

	if currentUnit.numIngreds > maxIngr || currentUnit.numMagimints > maxMags {
		return
	}
	magIdx, ok := getNeededMag(&currentUnit.desiredPotion, &currentUnit.magimints)
	if !ok {
		return
	}
	if magIdx == 5 && currentUnit.numIngreds >= minIngr && currentUnit.numMagimints >= minMags {
		sr := SearchResult{
			ResultingPotion: potionsMap[currentUnit.desiredPotion],
			TotalMagimints:  currentUnit.numMagimints,
			NumberIngreds:   currentUnit.numIngreds,
			Ingredients:     make(map[string]uint16, len(currentUnit.ingredients)),
		}
		for _, ingr := range currentUnit.ingredients {
			sr.Ingredients[ingr.name] = ingr.quantity
		}
		mu.Lock()
		*searchResult = append(*searchResult, sr)
		mu.Unlock()
	}

	numIngredsToAdd := maxIngr - currentUnit.numIngreds
	if numIngredsToAdd == 0 {
		return
	}

	if magIdx == 5 {
		for i = 0; i < 4; i++ {
			if currentUnit.desiredPotion[i] != 0 {
				break
			}
		}
		magIdx = i
	}
	for i = currentUnit.minIngredToUse[magIdx]; int(i) < len(IngredsByMags[magIdx]); i++ {
		ingrInfo := IngredsByMags[magIdx][i]
		ingrName := ingrInfo.name
		for _, ingr := range currentUnit.ingredients {
			if ingr.name == ingrName {
				goto end
			}
		}
		for j = 1; j <= numIngredsToAdd; j++ {
			newUnit := searchUnit{
				desiredPotion:  currentUnit.desiredPotion,
				magimints:      currentUnit.magimints,
				numIngreds:     currentUnit.numIngreds,
				numMagimints:   currentUnit.numMagimints,
				minIngredToUse: currentUnit.minIngredToUse,
				ingredients:    currentUnit.ingredients,
			}
			newUnit.minIngredToUse[magIdx] = i
			addIngred(&ingrInfo, &newUnit, j)

			wg.Add(1)
			stackMu.Lock()
			stack = append(stack, newUnit)
			stackMu.Unlock()
		}
	end:
	}
}

func addIngred(ingredInfo *NameWithMags, addTo *searchUnit, quantity uint16) {
	for i := range 5 {
		addTo.magimints[i] += ingredInfo.mags[i] * quantity
		addTo.numMagimints += ingredInfo.mags[i] * quantity
	}
	addTo.ingredients = append(addTo.ingredients, nameWithQuantity{ingredInfo.name, quantity})
	addTo.numIngreds += quantity
}

func getNeededMag(desiredPotion *[5]uint16, magimints *[5]uint16) (uint16, bool) {
	var maxMagIdx, maxMagValue uint16 = 100, 0
	for i, mag := range magimints {
		if mag > 0 && (*desiredPotion)[i] == 0 {
			return 0, false
		}
		if mag > maxMagValue {
			maxMagIdx = uint16(i)
			maxMagValue = mag
		}
	}
	if maxMagIdx == 100 {
		for i, val := range desiredPotion {
			if val != 0 {
				return uint16(i), true
			}
		}
	}
	delim := desiredPotion[maxMagIdx]
	if magimints[maxMagIdx]%delim != 0 {
		return maxMagIdx, true
	}
	unit := magimints[maxMagIdx] / delim
	for i, mag := range magimints {
		if desiredPotion[i]*unit > mag {
			return uint16(i), true
		} else if desiredPotion[i]*unit < mag {
			return maxMagIdx, true
		}
	}
	return 5, true
}

type searchUnit struct {
	desiredPotion [5]uint16
	magimints     [5]uint16
	numIngreds    uint16
	numMagimints  uint16
	// using map here is better by time complexity, but actually it's slower for reasonable number of ingreds
	ingredients    []nameWithQuantity
	minIngredToUse [5]uint16
}

type nameWithQuantity struct {
	name     string
	quantity uint16
}

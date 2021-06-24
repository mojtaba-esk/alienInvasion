// This package manges the aliens
// supported actions: Creating a new alien, Init mass invasion and individual invasion
package aliens

import (
	"alienInvasion/tools"
	"alienInvasion/world"
	"fmt"
	"sync"
)

// This type refers to a hashmap that holds all mutex of the cities
// By having the pointer to each city, we can quickly find the corresponding mutex
type citiesMutexList map[*world.City]*sync.Mutex

/**
* This function initiates the invasion.
* It first generates a list of aliens, then it put each alien in a randomly chosen city,
*  then run the Invade function of the alien in a separate goroutine
 */
func InitInvasion(numOfAliens int64, listOfCities world.Cities) {

	numOfCities := int64(len(listOfCities)) // better to cast once that having it in the loop

	if numOfAliens <= 0 || numOfCities == 0 {
		return
	}

	/*----------*/

	// This mutexList helps us to put exclusive locks on an individual city
	mutexList := make(citiesMutexList, numOfCities)

	//Init all mutex
	for i := int64(0); i < numOfCities; i++ {
		mutexList[listOfCities[i]] = &sync.Mutex{}
	}

	/*----------*/

	var wg sync.WaitGroup
	for i := int64(0); i < numOfAliens; i++ {

		// Generate a new alien
		name := fmt.Sprintf("A#%d", i)
		alien := newAlien(name)

		// Find a random city to put the alien in
		randomCityIndex := tools.RandomNumberI(0, numOfCities-1)
		randomCity := listOfCities[randomCityIndex]

		wg.Add(1)

		// Init the invasion, yoohooohahahhha >D
		go alien.Invade(randomCity, &wg, mutexList)

	}
	wg.Wait()
}

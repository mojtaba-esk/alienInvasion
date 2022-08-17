package aliens

import (
	"fmt"
	"sync"

	"github.com/mojtaba-esk/alienInvasion/tools"
	"github.com/mojtaba-esk/alienInvasion/world"
)

// This function recives the name of an alien,
// create a new alien object and returns the pointer address to it
func newAlien(name string) *Alien {
	return &Alien{name: name}
}

// This method start the invasion of the alien
// It first enters to the given city then finds all available outgoing paths of the city
// then it randomly chooses a path and moves to another city
// If it gets trapped i.e. there is no way to go out of the city, it terminates itself
// When an alien enters a city, we need to keep track of all invaders; in order to avoid race condition,
// we consider it as a critical section and protect each city with an individual mutex
func (a *Alien) Invade(currentCity *world.City, wg *sync.WaitGroup, mutexList sync.Map) {

	defer wg.Done()

	numOfVisits := 0
	for numOfVisits < _alienMoveThreshold {

		/*-------*/

		mxIf, ok := mutexList.Load(currentCity)
		if !ok {
			panic(fmt.Sprintf("Could not find the mutex for the city: %s", currentCity.GetName()))
		}
		mx, ok := mxIf.(*sync.Mutex)
		if !ok {
			panic(fmt.Sprintf("Could not convert the interface mutex for the city: %s", currentCity.GetName()))
		}
		mx.Lock()

		// Alien enters the city
		currentCity.Enter(a.name)

		mx.Unlock()

		/*-------*/

		// Check if there is any available path to go to the neighboring cities
		totalAvailablePaths := world.Cities{}

		if currentCity.GetNorth() != nil {
			totalAvailablePaths = append(totalAvailablePaths, currentCity.GetNorth())
		}
		if currentCity.GetEast() != nil {
			totalAvailablePaths = append(totalAvailablePaths, currentCity.GetEast())

		}
		if currentCity.GetSouth() != nil {
			totalAvailablePaths = append(totalAvailablePaths, currentCity.GetSouth())
		}
		if currentCity.GetWest() != nil {
			totalAvailablePaths = append(totalAvailablePaths, currentCity.GetWest())
		}

		if len(totalAvailablePaths) == 0 {
			// The city is destroyed it is the end of my life
			return
		}

		randomPathIndex := tools.RandomNumberI(0, int64(len(totalAvailablePaths))-1)

		/*-----------*/

		mx.Lock()

		//Leaving the current city
		currentCity.Leave(a.name)

		mx.Unlock()

		/*-----------*/

		currentCity = totalAvailablePaths[randomPathIndex]
		numOfVisits++

	}
}

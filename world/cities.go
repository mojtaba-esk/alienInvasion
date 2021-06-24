// This package manages the world
// Supported actions: handling the cities lists, generating new world, load database of cities
package world

import "fmt"

/*-----------*/

type Cities []*City

// This function receives a slice of cities and prints them all,
func PrintAllCities(listOfCities Cities) {

	numOfCities := len(listOfCities)

	for i := 0; i < numOfCities; i++ {
		cityStr := listOfCities[i].ToString()
		if cityStr != "" {
			fmt.Println(cityStr)
		}
	}
}

// This function counts and returns the number of destroyed cities
func NumOfDestroyedCities(listOfCities Cities) int64 {

	var output int64 = 0

	numOfCities := len(listOfCities)

	for i := 0; i < numOfCities; i++ {
		if listOfCities[i].Destroyed {
			output++
		}
	}
	return output
}

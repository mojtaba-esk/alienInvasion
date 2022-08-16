// This package manages the world
// Supported actions: handling the cities lists, generating new world, load database of cities
package world

import "fmt"

// This function counts and returns the number of destroyed cities
func GetNumOfDestroyedCities(listOfCities Cities) int64 {

	var output int64 = 0

	numOfCities := len(listOfCities)

	for i := 0; i < numOfCities; i++ {
		if listOfCities[i].destroyed {
			output++
		}
	}
	return output
}

// This function receives a slice of cities and prints them all,
func PrintAllCities(listOfCities Cities) {

	numOfCities := len(listOfCities)

	for i := 0; i < numOfCities; i++ {
		cityStr := listOfCities[i].String()
		if cityStr != "" {
			fmt.Println(cityStr)
		}
	}
}

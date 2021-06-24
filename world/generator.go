package world

import (
	"alienInvasion/tools"
	"bufio"
	"fmt"
	"log"
	"os"
)

/*--------------------*/

/**
* This function receives the number of cities and generates a new world accordingly
* It returns a slice holding all the pointers to the cities
 */
func Generate(numOfCities int64) Cities {

	if numOfCities <= 0 {
		return Cities{}
	}

	listOfAllCities := make(Cities, numOfCities)

	// Generate new cities
	for i := int64(0); i != numOfCities; i++ {
		listOfAllCities[i] = generateNewCity()
	}

	// Connect the cities to eachother randomly
	for i := int64(0); i != numOfCities-1; i++ {
		listOfAllCities[i].randomConnect(listOfAllCities[i+1:])
	}

	return listOfAllCities

}

// `cityIdSequence` is a global var that is used to generate sequential names
// for cities in case the real names are not available or not enough
var cityIdSequence int64

/**
* This function generates a new random city and returns a pointer to its node
 */
func generateNewCity() *City {

	var cityNode City

	randomCityName := getRandomCityName()
	if randomCityName == "" {
		randomCityName = fmt.Sprintf("City_#%d", cityIdSequence)
		cityIdSequence++
	}
	cityNode.Name = randomCityName

	return &cityNode
}

/**
* This function gets a random unique name from the city database
 */
var cityNamesList []string

func getRandomCityName() string {

	// Load the city names for the first time
	if cityNamesList == nil {
		var err error
		cityNamesList, err = loadCityNamesFromCSV()

		if err != nil {
			log.Printf("[Err  ] city names .csv file loading: %v", err)
			return ""
		}
	}

	if len(cityNamesList) == 0 {
		return ""
	}

	// Find a unique random name
	randomIndex := tools.RandomNumberI(0, int64(len(cityNamesList))-1)
	output := cityNamesList[randomIndex]

	// Let's remove the city name in order to avoid duplicates
	cityNamesList = tools.SliceItemRemoveByIndex(cityNamesList, randomIndex)

	return output
}

/**
* This function reads the csv file locates in `/data/world-cities.csv` and
* returns a slice with the names of all the major cities in the world
 */

func loadCityNamesFromCSV() ([]string, error) {

	file, err := os.Open(cityNamesCSVFilePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil

}

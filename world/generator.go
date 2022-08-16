package world

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mojtaba-esk/alienInvasion/tools"
)

var (
	// `_cityIdSequence` is a global var that is used to generate sequential names
	// for cities in case the real names are not available or not enough
	_cityIdSequence int64

	// This function gets a random unique name from the city database
	_cityNamesList []string
)

// This function receives the number of cities and generates a new world accordingly
// It returns a slice holding all the pointers to the cities
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

// This function generates a new random city and returns a pointer to its node
func generateNewCity() *City {

	var cityNode City

	randomCityName := getRandomCityName()
	if randomCityName == "" {
		randomCityName = fmt.Sprintf("City_#%d", _cityIdSequence)
		_cityIdSequence++
	}
	cityNode.name = randomCityName

	return &cityNode
}

func getRandomCityName() string {

	// Load the city names for the first time
	if _cityNamesList == nil {
		var err error
		_cityNamesList, err = loadCityNamesFromCSV()

		if err != nil {
			log.Printf("[Err  ] city names .csv file loading: %v", err)
			return ""
		}
	}

	if len(_cityNamesList) == 0 {
		return ""
	}

	// Find a unique random name
	randomIndex := tools.RandomNumberI(0, int64(len(_cityNamesList))-1)
	output := _cityNamesList[randomIndex]

	// Let's remove the city name in order to avoid duplicates
	_cityNamesList = tools.SliceItemRemoveByIndex(_cityNamesList, randomIndex)

	return output
}

// This function reads the csv file locates in `/data/world-cities.csv` and
// returns a slice with the names of all the major cities in the world
func loadCityNamesFromCSV() ([]string, error) {

	file, err := os.Open(_cityNamesCSVFilePath)
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

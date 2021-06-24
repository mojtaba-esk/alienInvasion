package world

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

/*---------------------*/

// This function receives a `filepath`, attempts to load the file
// then hands over the content to `ParseText()` for further processing
// It returns the parsed data in form of a slice of pointers to all cities
func Load(filepath string) (Cities, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			return
		}
	}()

	contentInBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	text := string(contentInBytes)
	return ParseText(text)
}

// This function receives a `text` and parses the data to extract city names and their connection
// It returns the parsed data in form of a slice of pointers to all cities
func ParseText(text string) (Cities, error) {

	var listOfCities Cities

	hashTable := make(map[string]*City)

	lines := strings.Split(text, "\n")
	// Extracting the data rows
	for _, line := range lines {

		if line == "" {
			continue
		}

		line := strings.Trim(line, " \t\n\r")

		reCityName := regexp.MustCompile(`^[^\ ]+`)
		cityName := reCityName.FindString(line)

		// Check if we have added the current city before
		currentCity := hashTable[cityName]

		// If it is new, let's add it to our hashTable
		if currentCity == nil {
			hashTable[cityName] = &City{Name: cityName}
			currentCity = hashTable[cityName]
		}

		// Handling the neighboring cities
		reConnectedCities := regexp.MustCompile(`\ (north|east|south|west)=([^\ ]+)`)
		connectedMatches := reConnectedCities.FindAllStringSubmatch(line, -1)

		for _, match := range connectedMatches {

			targetDirection := match[1]
			targetCityName := match[2]

			// Check if we have seen the target city before
			targetCity := hashTable[targetCityName]

			// If it is new, let's add it to our hashTable
			if targetCity == nil {
				hashTable[targetCityName] = &City{Name: targetCityName}
				targetCity = hashTable[targetCityName]
			}

			// Handling the direction
			switch targetDirection {
			case "north":
				currentCity.North = targetCity
			case "east":
				currentCity.East = targetCity
			case "south":
				currentCity.South = targetCity
			case "west":
				currentCity.West = targetCity
			default:
				return nil, fmt.Errorf("malformatted data: direction is not corrent")
			}

		}

		listOfCities = append(listOfCities, currentCity)
	}

	return listOfCities, nil
}

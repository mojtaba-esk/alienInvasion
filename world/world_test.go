package world

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// Disabling logs
func TestMain(m *testing.M) {

	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

/**
* This function tests the world (cities) generator
* and data load and compares the results to see
* if everything regarding data-generation, data-load, and world graph construction work well
 */
func TestGenerateAndLoad(t *testing.T) {

	// Test table
	tt := []struct {
		name        string
		numOfCities int64 // Number of cities
	}{
		{
			name:        "Testing with large number of cities",
			numOfCities: 5000,
		},
		{
			name:        "Testing with avarage number of cities",
			numOfCities: 100,
		},
		{
			name:        "Testing with small number of cities",
			numOfCities: 5,
		},
		{
			name:        "Testing with 1",
			numOfCities: 1,
		},
		{
			name:        "Testing with Zero",
			numOfCities: 0,
		},
		{
			name:        "Testing with a negative number",
			numOfCities: -5,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			listOfCities := Generate(tc.numOfCities)

			if tc.numOfCities >= 0 && len(listOfCities) != int(tc.numOfCities) {
				t.Fatalf("Expected to generate `%d` cities, but get `%d` cities", tc.numOfCities, len(listOfCities))
			}

			// Building the database
			database := buildTextDatabase(listOfCities)

			// Testing Database Load
			loadedCities, err := ParseText(database)
			if err != nil {
				t.Fatalf("Error in database load: %v", err)
			}

			newDatabase := buildTextDatabase(loadedCities)
			if newDatabase != database {
				t.Fatalf("Error in database load. A mismatch detected! Expected db length: `%d`, loaded db length: `%d`", len(database), len(newDatabase))
			}

		})
	}
}

/*---------------*/

func buildTextDatabase(listOfCities Cities) string {

	output := ""
	for i := 0; i < len(listOfCities); i++ {
		cityStr := listOfCities[i].ToString()
		if cityStr != "" {
			output += cityStr + "\n"
		}
	}

	return output
}

package aliens

import (
	"alienInvasion/world"
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

// This function tests the aliens package
// it first generates a new world and then launches various invasions on it
func TestAliens(t *testing.T) {

	// Test table
	tt := []struct {
		name         string
		numOfCities  int64 // Number of cities
		numOfAliens  int64 // Number of aliens
		minDestroyed int64 // Minimum expected number of destroyed cities
		maxDestroyed int64 // Maximum expected number of destroyed cities
	}{
		{
			name:         "Testing 5000 cities 1000 aliens",
			numOfCities:  5000,
			numOfAliens:  1000,
			minDestroyed: 10,
			maxDestroyed: 5000,
		},
		{
			name:         "Testing 500 cities 10 aliens",
			numOfCities:  500,
			numOfAliens:  10,
			minDestroyed: 0,
			maxDestroyed: 500,
		},
		{
			name:         "Testing 50 cities 10 aliens",
			numOfCities:  50,
			numOfAliens:  10,
			minDestroyed: 0,
			maxDestroyed: 50,
		},
		{
			name:         "Testing 50 cities 100 aliens",
			numOfCities:  50,
			numOfAliens:  100,
			minDestroyed: 1,
			maxDestroyed: 50,
		},
		{
			name:         "Testing 50 cities 1000 aliens",
			numOfCities:  50,
			numOfAliens:  1000,
			minDestroyed: 40, // rather 50
			maxDestroyed: 50,
		},
		{
			name:         "Testing 50 cities 2 aliens",
			numOfCities:  50,
			numOfAliens:  2,
			minDestroyed: 0,
			maxDestroyed: 1,
		},
		{
			name:         "Testing 50 cities 1 aliens",
			numOfCities:  50,
			numOfAliens:  1,
			minDestroyed: 0,
			maxDestroyed: 0,
		},
		{
			name:         "Testing 50 cities 0 aliens",
			numOfCities:  50,
			numOfAliens:  0,
			minDestroyed: 0,
			maxDestroyed: 0,
		},
		{
			name:         "Testing 50 cities -10 aliens",
			numOfCities:  50,
			numOfAliens:  -10,
			minDestroyed: 0,
			maxDestroyed: 0,
		},
		{
			name:         "Testing 0 cities 10 aliens",
			numOfCities:  0,
			numOfAliens:  10,
			minDestroyed: 0,
			maxDestroyed: 0,
		},
		{
			name:         "Testing -5 cities 10 aliens",
			numOfCities:  -5,
			numOfAliens:  10,
			minDestroyed: 0,
			maxDestroyed: 0,
		},
		{
			name:         "Testing -5 cities -10 aliens",
			numOfCities:  -5,
			numOfAliens:  -10,
			minDestroyed: 0,
			maxDestroyed: 0,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			listOfCities := world.Generate(tc.numOfCities)

			InitInvasion(tc.numOfAliens, listOfCities)

			numOfDestroyed := world.NumOfDestroyedCities(listOfCities)

			if numOfDestroyed < tc.minDestroyed {
				t.Fatalf("Expected to destory at least `%d` cities, but `%d` cities destroyed", tc.minDestroyed, numOfDestroyed)
			}
			if numOfDestroyed > tc.maxDestroyed {
				t.Fatalf("Expected to destory max `%d` cities, but `%d` cities destroyed", tc.maxDestroyed, numOfDestroyed)
			}

		})
	}
}

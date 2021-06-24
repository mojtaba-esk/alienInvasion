package world

import (
	"alienInvasion/tools"
	"fmt"
)

/*-----------*/

type City struct {
	Name string

	North *City
	East  *City
	South *City
	West  *City

	InvaderNames []string // We need only the names of invaders
	Destroyed    bool     // A flag indicating if the city is already destroyed (default:false)
}

/*-----------*/

// This function connects the city to a set of randomly selected cities from the given city list to random directions
func (c *City) randomConnect(listOfCities Cities) {

	// Generate a random number for each direction to decide to connect it or not
	randomDirections := make([]int64, 4)

	// Total random directions for this particular node (e.g. north, south, west = 3 directions)
	// This is specially useful for task division
	totalDirections := int64(0) // int64: just to avoid too many type casting

	// Generating random direction probability (50% chance for every direction)
	for i := 0; i != 4; i++ {
		randomDirections[i] = tools.RandomNumberI(0, 1)
		totalDirections += randomDirections[i]
	}
	/*-----------*/

	// If no direction get a chance, we select one direction randomly in order to avoid isolated cities
	if totalDirections == 0 {
		randomDirections[tools.RandomNumberI(0, 3)] = 1
		totalDirections = 1
	}

	/*-----------*/

	nextCityIndex := 0

	if c.North == nil && randomDirections[0] == 1 && nextCityIndex < len(listOfCities) {
		c.North = listOfCities[nextCityIndex]
		listOfCities[nextCityIndex].South = c
		nextCityIndex++
	}

	if c.East == nil && randomDirections[1] == 1 && nextCityIndex < len(listOfCities) {
		c.East = listOfCities[nextCityIndex]
		listOfCities[nextCityIndex].West = c
		nextCityIndex++
	}

	if c.South == nil && randomDirections[2] == 1 && nextCityIndex < len(listOfCities) {
		c.South = listOfCities[nextCityIndex]
		listOfCities[nextCityIndex].North = c
		nextCityIndex++
	}

	if c.West == nil && randomDirections[3] == 1 && nextCityIndex < len(listOfCities) {
		c.West = listOfCities[nextCityIndex]
		listOfCities[nextCityIndex].East = c
		nextCityIndex++
	}

}

// This function is called once the city is visited by an alien
// The invasion result can happen here
func (c *City) Enter(invaderName string) {

	c.InvaderNames = append(c.InvaderNames, invaderName)

	if !c.Destroyed && len(c.InvaderNames) >= 2 {

		// Let's destroy the city
		c.printDestructionMessage()
		c.destroy()
	}

}

// This function is called once the city is left by an alien
func (c *City) Leave(invaderName string) {

	// We just remove the name of the alien from the invaders
	c.InvaderNames = tools.SliceItemRemove(c.InvaderNames, invaderName)
}

// This function prepares the printable string of the available city (i.e. if not destroyed)
// in the same format of the database
func (c *City) ToString() string {

	output := ""
	if c == nil || c.Destroyed {
		return output
	}

	output += fmt.Sprintf("%s ", c.Name)

	if c.North != nil && !c.North.Destroyed {
		output += fmt.Sprintf("north=%s ", c.North.Name)
	}

	if c.East != nil && !c.East.Destroyed {
		output += fmt.Sprintf("east=%s ", c.East.Name)
	}

	if c.South != nil && !c.South.Destroyed {
		output += fmt.Sprintf("south=%s ", c.South.Name)
	}

	if c.West != nil && !c.West.Destroyed {
		output += fmt.Sprintf("west=%s ", c.West.Name)
	}

	return output
}

// This function prints the fight message and destruction report of the city
func (c *City) printDestructionMessage() {

	if len(c.InvaderNames) == 0 {
		return
	}

	fmt.Printf("\n`%s` has been destroyed by ", c.Name)
	for i, alienName := range c.InvaderNames {

		if i == 0 {
			fmt.Printf("alien `%s`", alienName)
		} else {
			fmt.Printf(" and alien `%s`", alienName)
		}
	}
	fmt.Printf("!\n")

}

// This function destroyes a city according to the given instructions
func (c *City) destroy() {

	// First Block the path from the neighbors to this city
	if c.North != nil {
		c.North.South = nil
	}
	if c.East != nil {
		c.East.West = nil
	}
	if c.South != nil {
		c.South.North = nil
	}
	if c.West != nil {
		c.West.East = nil
	}

	// Clean up the city
	c.North = nil
	c.East = nil
	c.South = nil
	c.West = nil

	c.Name = "" // just to save some space ;)
	c.Destroyed = true
}

/*----------*/

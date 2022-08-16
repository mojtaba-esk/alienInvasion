package world

import (
	"fmt"

	"github.com/mojtaba-esk/alienInvasion/tools"
)

// This function is called once the city is visited by an alien
// The invasion result can happen here
func (c *City) Enter(invaderName string) {

	c.invaderNames = append(c.invaderNames, invaderName)

	if !c.destroyed && len(c.invaderNames) >= 2 {

		// Let's destroy the city
		c.printDestructionMessage()
		c.destroy()
	}

}

// This function is called once the city is left by an alien
func (c *City) Leave(invaderName string) {

	// We just remove the name of the alien from the invaders
	c.invaderNames = tools.SliceItemRemove(c.invaderNames, invaderName)
}

func (c *City) GetName() string   { return c.name }
func (c *City) GetNorth() *City   { return c.north }
func (c *City) GetEast() *City    { return c.east }
func (c *City) GetSouth() *City   { return c.south }
func (c *City) GetWest() *City    { return c.west }
func (c *City) IsDestroyed() bool { return c.destroyed }

func (c *City) GetInvaderNames() []string {
	output := make([]string, len(c.invaderNames))
	copy(output, c.invaderNames)
	return output
}

// This function prepares the printable string of the available city (i.e. if not destroyed)
// in the same format of the database
func (c *City) String() string {

	output := ""
	if c == nil || c.destroyed {
		return output
	}

	output += fmt.Sprintf("%s ", c.name)

	if c.north != nil && !c.north.destroyed {
		output += fmt.Sprintf("north=%s ", c.north.name)
	}

	if c.east != nil && !c.east.destroyed {
		output += fmt.Sprintf("east=%s ", c.east.name)
	}

	if c.south != nil && !c.south.destroyed {
		output += fmt.Sprintf("south=%s ", c.south.name)
	}

	if c.west != nil && !c.west.destroyed {
		output += fmt.Sprintf("west=%s ", c.west.name)
	}

	return output
}

// This function destroys a city according to the given instructions
func (c *City) destroy() {

	// First Block the path from the neighbors to this city
	if c.north != nil {
		c.north.south = nil
	}
	if c.east != nil {
		c.east.west = nil
	}
	if c.south != nil {
		c.south.north = nil
	}
	if c.west != nil {
		c.west.east = nil
	}

	// Clean up the city
	c.north = nil
	c.east = nil
	c.south = nil
	c.west = nil

	c.name = "" // just to save some space ;)
	c.destroyed = true
}

/*----------*/

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

	if c.north == nil && randomDirections[0] == 1 && nextCityIndex < len(listOfCities) {
		c.north = listOfCities[nextCityIndex]
		listOfCities[nextCityIndex].south = c
		nextCityIndex++
	}

	if c.east == nil && randomDirections[1] == 1 && nextCityIndex < len(listOfCities) {
		c.east = listOfCities[nextCityIndex]
		listOfCities[nextCityIndex].west = c
		nextCityIndex++
	}

	if c.south == nil && randomDirections[2] == 1 && nextCityIndex < len(listOfCities) {
		c.south = listOfCities[nextCityIndex]
		listOfCities[nextCityIndex].north = c
		nextCityIndex++
	}

	if c.west == nil && randomDirections[3] == 1 && nextCityIndex < len(listOfCities) {
		c.west = listOfCities[nextCityIndex]
		listOfCities[nextCityIndex].east = c
		nextCityIndex++
	}

}

// This function prints the fight message and destruction report of the city
func (c *City) printDestructionMessage() {

	if len(c.invaderNames) == 0 {
		return
	}

	fmt.Printf("\n`%s` has been destroyed by ", c.name)
	for i, alienName := range c.invaderNames {

		if i == 0 {
			fmt.Printf("alien `%s`", alienName)
		} else {
			fmt.Printf(" and alien `%s`", alienName)
		}
	}
	fmt.Printf("!\n")

}

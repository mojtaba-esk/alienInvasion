package world

/*-----------*/

type City struct {
	name string

	north *City
	east  *City
	south *City
	west  *City

	invaderNames []string // We need only the names of invaders
	destroyed    bool     // A flag indicating if the city is already destroyed (default:false)
}

type Cities []*City

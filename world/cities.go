// This package manages the world
// Supported actions: handling the cities lists, generating new world, load database of cities
package world

/*-----------*/

type Cities []*City

/**
*
* This function receives a slice of cities and prints them all,
*
 */

func PrintAllCities(listOfCities Cities) {

	numOfCities := len(listOfCities)

	for i := 0; i < numOfCities; i++ {
		listOfCities[i].print()
	}
}

/**
* This function receives a slice of cities and prints them all,
* useful for storing the generated cities in a file and for debugging
 */

// func CheckAllCities(listOfCities Cities) {

// 	numOfCities := len(listOfCities)

// 	fmt.Printf("numOfCities: %+v\n", numOfCities)

// 	// Print all cities
// 	for i := 0; i < numOfCities; i++ {
// 		fmt.Printf("\nCity `%s`", listOfCities[i].Name)

// 		// nextCity := listOfCities[i].North
// 		if listOfCities[i].North != nil {
// 			fmt.Printf("\n\tNorth: %p => `%s`", listOfCities[i].North, listOfCities[i].North.Name)

// 			found := false
// 			// Search for the next city in the list
// 			for j := 0; j < numOfCities; j++ {
// 				if listOfCities[i].North == listOfCities[j] {
// 					fmt.Printf("\tFound in the list: %p => `%s`", listOfCities[j], listOfCities[j].Name)
// 					found = true
// 					break
// 				}
// 			}

// 			if !found {
// 				fmt.Printf("\tNot Found!\n\n")
// 				return
// 			}

// 		}

// 	}
// }

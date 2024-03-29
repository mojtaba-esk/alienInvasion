package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mojtaba-esk/alienInvasion/aliens"
	"github.com/mojtaba-esk/alienInvasion/world"
)

/*--------------*/

func main() {

	configLogs()

	/*--------------*/

	var (
		action       string
		numOfCities  int64
		databasePath string
		numOfAliens  int64
	)

	flag.StringVar(&action, "a", "start", `Specify the action: 
	- generate-world	Generates a new world
	- start			Starts the invasion normally`,
	)
	flag.Int64Var(&numOfCities, "c", 10, "Specify the number of cities for the world generator.")
	flag.StringVar(&databasePath, "d", "database.txt", "Specify the path to the database file.")
	flag.Int64Var(&numOfAliens, "n", 5, "Specify the number of aliens.")

	flag.Parse()

	/*--------------*/

	switch action {
	case "start":
		{
			listOfAllCities, err := world.Load(databasePath)
			if err != nil {
				log.Printf("[Err  ] database loading: %v", err)
				panic(err)
			}
			fmt.Printf("\n%d cities are loaded.", len(listOfAllCities))
			fmt.Printf("\nInitiating the invasion with %d aliens...", numOfAliens)

			fmt.Printf("\n\n\t\t==============================\n")
			fmt.Printf("\nThe invasion report:\n")

			aliens.InitInvasion(numOfAliens, listOfAllCities)

			fmt.Printf("\n\n\t\t==============================\n")

			totalAliveCities := int64(len(listOfAllCities)) - world.GetNumOfDestroyedCities(listOfAllCities)
			fmt.Printf("\nThe cities left of the world: `%d`\n\n", totalAliveCities)

			world.PrintAllCities(listOfAllCities)
		}
	case "generate-world":
		{
			world.PrintAllCities(world.Generate(numOfCities))
		}
	default:
		{
			fmt.Println("Wrong action. ")
			flag.PrintDefaults()
		}
	}

}

/*--------------*/

//  This function configures the output of the logs
func configLogs() {

	// Let's store logs in a log file
	filePtr, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(filePtr)
}

// A project to simulate traffic navigation in the city of Hastings.
//
// Team members: Frederick Stoddart, Levi Fraser-Daley, Nikolas Burke
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// A location representing a suburb in a city
type funCityLoc struct {
	Name       string
	Connectors []funConnector
}

// A collection of streets that connect two areas and an exit to leave the city
type funConnector struct {
	Destinations []string
	Streets      []funStreet
	Exit         funExit
}

// A street
type funStreet string

// An exit leading to Outside City
type funExit string

// A driver to be used in the simulation
type funDriver struct {
	Number int
	Name   string
}

// Setting up the state for our simulation
var (
	akinaStortfordLodgeConnector = funConnector{
		[]string{
			"Akina",
			"Stortford Lodge",
		},
		[]funStreet{
			"Southhampton St W",
		},
		"Railway Rd S",
	}

	stortfordLodgeMahoraConnector = funConnector{
		[]string{
			"Mahora",
		},
		[]funStreet{
			"Maraekakaho Rd",
			"Pakowhai Rd",
		},
		"Omahu Rd",
	}

	mahoraMayfairConnector = funConnector{
		[]string{
			"Mayfair",
		},
		[]funStreet{
			"Frederick St",
			"Grove Rd",
		},
		"Karamu Rd",
	}

	akinaMayfairConnector = funConnector{
		[]string{
			"Akina",
			"Mayfair",
		},
		[]funStreet{
			"Willowpark Rd",
		},
		"Havelock Rd",
	}

	hastings = []funCityLoc{
		{
			"Mayfair",
			[]funConnector{
				akinaMayfairConnector,
			},
		},
		{
			"Akina",
			[]funConnector{
				akinaMayfairConnector,
				akinaStortfordLodgeConnector,
			},
		},
		{
			"Stortford Lodge",
			[]funConnector{
				stortfordLodgeMahoraConnector,
				akinaStortfordLodgeConnector,
			},
		},
		{
			"Mahora",
			[]funConnector{
				mahoraMayfairConnector,
			},
		},
	}

	drivers = []funDriver{
		{1, "Fred"},
		{2, "Caitlyn"},
		{3, "Mason"},
		{4, "Bea"},
		{5, "Tara"},
	}
)

// Main() :p
func main() {


	args := os.Args
	// TODO check args are correct

	seed, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	rand.Seed(seed)

	for i := 0; i < len(drivers); i++ {

		log := fmt.Sprintf("Driver: %v, is starting their trip \n", drivers[i].Name)

		s := rand.NewSource(rand.Int63())

		r := rand.New(s)

		currentPos := startSim(r)

		log = log + fmt.Sprintf("They are starting in %v \n", currentPos.Name)

		insideCity := true

		for insideCity {

			// TODO continue with inside city code

			route := pickRoute(r, currentPos)
			j, err := json.Marshal(&currentPos)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(string(j))

			if randInRange(0, 6, r) == 0 {
				insideCity = false
				fmt.Println(log)
			} else {
				if len(route.Destinations) == 1 {
					currentPos = getLoc(route.Destinations[0])
				} else {
					if getLoc(route.Destinations[0]).Name == currentPos.Name {
						currentPos = getLoc(route.Destinations[1])
					} else {
						currentPos = getLoc(route.Destinations[0])
					}
				}
			}
		}
	}

}

func startSim(rand *rand.Rand) *funCityLoc {
	return &hastings[randInRange(0, int64(len(hastings)), rand)]
}

func randInRange(min, max int64, rand *rand.Rand) int64 {
	return rand.Int63n(max-min) + min
}

func pickRoute(rand *rand.Rand, currentPos *funCityLoc) *funConnector {
	c := &currentPos.Connectors[randInRange(0, int64(len(currentPos.Connectors)), rand)]
	return c
}

func getLoc(locName string) *funCityLoc {
	for i := 0; i < len(hastings); i++ {
		if locName == hastings[i].Name {
			return &hastings[i]
		}
	}
	return nil
}
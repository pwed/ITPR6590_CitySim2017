// A project to simulate traffic navigation in the city of Hastings.
//
// Team members: Frederick Stoddart, Levi Fraser-Daley, Nikolas Burke
package main

import (
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
		{5, "Hayley"},
	}
)

// Main() :p
func main() {

	// Fun Args
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Too many or not enough arguments!")
		os.Exit(1)
	}

	seed, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	rand.Seed(seed)

	for i := 0; i < len(drivers); i++ {

		log := fmt.Sprintf("Driver %v, is beginning their trip, ", drivers[i].Name)

		s := rand.NewSource(rand.Int63())

		r := rand.New(s)

		akinaCount := 0

		currentPos := startSim(r)

		prevPos := currentPos

		log = log + fmt.Sprintf("they are starting in %v. \n", currentPos.Name)

		insideCity := true

		for insideCity {

			// TODO continue with inside city code

			route := pickRoute(r, currentPos)

			if randInRange(0, 6, r) == 0 {
				insideCity = false

				log = log + fmt.Sprintf("Driver %v heading %v city to Outside City via %v. \n",
					drivers[i].Name, currentPos.Name, route.Exit)

				// Fun Other Cities
				log = otherCities(route, log, drivers[i].Name)

				log = log + fmt.Sprintf("Driver %v visited John Jamerson in Akina %v times. \n",
					drivers[i].Name, akinaCount)

				if akinaCount >= 3 {
					log = log + fmt.Sprintf("Driver %v needed lots of help! \n",
						drivers[i].Name)
				}

				if akinaCount == 0 {
					log = log + fmt.Sprintf("Driver %v missed out! \n",
						drivers[i].Name)
				}


				// Fun Dashes
				log = log + fmt.Sprintf("-----")

				fmt.Println(log)
			} else {

				prevPos = currentPos

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
			if insideCity {
				log = log + fmt.Sprintf("Driver %v heading from %v to %v via ", drivers[i].Name,
					prevPos.Name, currentPos.Name)
				for i, street := range route.Streets {
					if i > 1 {
						log = log + fmt.Sprintf(" and ",)
					}
					log = log + fmt.Sprintf("%v", street,)
				}
				log = log + fmt.Sprintf(". \n",)
			}
			akinaCounting(currentPos, &akinaCount)
		}
	}

}

func otherCities(route *funConnector, log string, driver string) string {
	if route.Exit == "Karamu Rd" {
		return log + fmt.Sprintf("Driver %v left and gone to %v. \n", driver, "Napier")
	}
	if route.Exit == "Omahu Rd" {
		return log + fmt.Sprintf("Driver %v left and gone to %v. \n", driver, "Flaxmere")
	}
	return log
}

func akinaCounting(pos *funCityLoc, count *int) {
	if pos.Name == "Akina" {
		*count++
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

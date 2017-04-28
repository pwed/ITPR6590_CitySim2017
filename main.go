// A project to simulate traffic navigation in the city of Hastings.
//
// Team members: Frederick Stoddart, Levi Fraser-Daley, Nikolas Burke
package main

import (
	"encoding/json"
	"fmt"
)

// A location representing a suburb in a city
type funCityLoc struct {
	Name       string
	Connectors []funConnector
}

// A collection of streets that connect two areas and an exit to leave the city
type funConnector struct {
	Streets []funStreet
	Exit    funExit
}

// A street connecting two locations and an exit.
// Can comprise of multiple physical streets
type funStreet string

// An exit leading to Outside City
type funExit string

// A driver to be used in the simulation
type funDriver struct {
	Number int
	Name   string
}

var akinaStortfordLodgeConnector = funConnector{
	[]funStreet{
		"Southhampton St W",
	},
	"Railway Rd S",
}

var stortfordLodgeMahoraConnector = funConnector{
	[]funStreet{
		"Maraekakaho Rd",
		"Pakowhai Rd",
	},
	"Omahu Rd",
}

var mahoraMayfairConnector = funConnector{
	[]funStreet{
		"Frederick St",
		"Grove Rd",
	},
	"Karamu Rd",
}

var akinaMayfairConnector = funConnector{
	[]funStreet{
		"Willowpark Rd",
	},
	"Havelock Rd",
}

var hastings = []funCityLoc{
	{
		"Mayfair",
		[]funConnector{
			akinaMayfairConnector,
			mahoraMayfairConnector,
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
			stortfordLodgeMahoraConnector,
		},
	},
}

var drivers = []funDriver{
	{1, "Fred"},
	{2, "Caitlyn"},
	{3, "Mason"},
	{4, "Bea"},
	{5, "Tara"},
}

// Main() :p
func main() {

	// This code is just to clear the unused variable warnings (and I made it output pretty JSON)
	hastingsJson, _ := json.MarshalIndent(hastings, "", "  ")
	fmt.Println(string(hastingsJson))
	fmt.Println(drivers)

}

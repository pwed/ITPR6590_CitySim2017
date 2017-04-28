// A project to simulate traffic navigation in the city of Hastings.
//
// Team members: Frederick Stoddart, Levi Fraser-Daley, Nikolas Burke
package main

import "fmt"

// A location representing a suburb in a city
type funCityLoc struct {
	name        string
	outsideCity bool
}

//
type funStreet struct {
	name string
	exit funExit
}

type funExit string

type funDriver struct {
	number int
	name   string
}

var hastings = []funCityLoc{
	{"Mayfair", false},
	{"Akina", false},
	{"Stortford Lodge", false},
	{"Mahora", false},
	{"Outside City", true},
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

	fmt.Println(hastings)
	fmt.Println(drivers)

}

// A project to simulate traffic navigation in the city of Hastings.
//
// Team members: Frederick Stoddart, Levi Fraser-Daley, Nikolas Burke
package main

import "fmt"


type funCityLoc struct {
	name string
	outsideCity bool
}

type funStreet struct {
	name string
	exit funExit
}

type funExit string

type funDriver struct {
	number int
	name string
}

var hastings = []funCityLoc{
	{"Mayfair", false},
	{"Akina", false},
	{"Stortford Lodge", false},
	{"Mahora", false},
	{"Outside City", true},
}

// Main() :p
func main() {

	fmt.Println(hastings)

}

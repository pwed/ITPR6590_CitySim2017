package main

import "testing"

func TestGetLoc(t *testing.T) {
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

	if getLoc(hastings[2].Name) != &hastings[2] {
		t.Error("Locations should be the same")
	}

	if getLoc(hastings[2].Name) == &hastings[1] {
		t.Error("Locations should be different")
	}

	if getLoc("Levi's house") != nil {
		t.Error("No location should be found, should return nil")
	}
}

func TestAkinaCountingt(t *testing.T) {
	count := 0

	notAkina := funCityLoc{
		"Mayfair",
		[]funConnector{
			akinaMayfairConnector,
		},
	}

	akina := funCityLoc{
		"Akina",
		[]funConnector{
			akinaMayfairConnector,
			akinaStortfordLodgeConnector,
		},
	}

	akinaCounting(&notAkina, &count)
	if count != 0 {t.Error("Count should not have changed")}

	akinaCounting(&akina, &count)
	if count != 1 {t.Error("Count should now be one")}
}

func TestOtherCities(t *testing.T) {
	omahu := funConnector{
		[]string{
			"Mahora",
		},
		[]funStreet{
			"Maraekakaho Rd",
			"Pakowhai Rd",
		},
		"Omahu Rd",
	}

	karamu := funConnector{
		[]string{
			"Mayfair",
		},
		[]funStreet{
			"Frederick St",
			"Grove Rd",
		},
		"Karamu Rd",
	}

	railway := funConnector{
		[]string{
			"Akina",
			"Stortford Lodge",
		},
		[]funStreet{
			"Southhampton St W",
		},
		"Railway Rd S",
	}
	fred := funDriver{1, "Fred"}

	log := ""

	if otherCities(&karamu, log, fred.Name) != "Driver Fred left and gone to Napier. \n" {
		t.Error("Strings should match")
	}

	if otherCities(&omahu, log, fred.Name) != "Driver Fred left and gone to Flaxmere. \n" {
		t.Error("Strings should match")
	}

	if otherCities(&railway, log, fred.Name) != "" {
		t.Error("Strings should be empty")
	}
}

func TestAkinaEdges(t *testing.T) {
	fred := funDriver{1, "Fred"}

	log := ""

	if akinaEdges(0, log, fred) != "Driver Fred missed out! \n" {
		t.Error("Fred should be missing out")
	}

	if akinaEdges(1, log, fred) != "" || akinaEdges(2, log, fred) != ""  {
		t.Error("an empty string should be returned")
	}

	if akinaEdges(3, log, fred) != "Driver Fred needed lots of help! \n" {
		t.Error("Fred should be needing help")
	}
}

func TestCheckArgs(t *testing.T) {

}

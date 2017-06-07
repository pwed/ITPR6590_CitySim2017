package main

import (
	"testing"
	"os"
	"os/exec"
	"reflect"
)

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

	tooManyArgs := []string{"appname", "1", "4"}

	argsNotAnInt := []string{"appname", "h"}

	goodArgs := []string{"appname", "1"}

	notEnoughArgs := []string{"appname"}

	if os.Getenv("too_many") == "1" {
		checkArgs(tooManyArgs)
		return
	}

	if os.Getenv("not_enough") == "1" {
		checkArgs(notEnoughArgs)
		return
	}

	if os.Getenv("not_int") == "1" {
		checkArgs(argsNotAnInt)
		return
	}

	if reflect.TypeOf(checkArgs(goodArgs)).String() != "int64" {
		t.Error("should return an int64")
	}

	// The following tests are not complete and need to be fixed.
	// The problem is that they call os.Exit() in the test and I'm not sure how to go about that

	cmd := exec.Command(os.Args[0], "-test.run=TestCheckArgs")
	cmd.Env = append(os.Environ(), "too_many=1")
	err := cmd.Run()
	if err == nil {
		t.Error("Process should exit with exit code 1")
	}

	cmd = exec.Command(os.Args[0], "-test.run=TestCheckArgs")
	cmd.Env = append(os.Environ(), "not_enough=1")
	err = cmd.Run()
	if err == nil {
		t.Error("Process should exit with exit code 1")
	}

	cmd = exec.Command(os.Args[0], "-test.run=TestCheckArgs")
	cmd.Env = append(os.Environ(), "not_int=1")
	err = cmd.Run()
	if err == nil {
		t.Error("Process should exit with exit code 2")
	}

}

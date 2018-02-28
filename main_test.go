package main

import (
	"fmt"
	"testing"

	geo "github.com/kellydunn/golang-geo"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestFunctions(t *testing.T) {
	customers, err := readLines(FilePath)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, locValidate(customers[0]), true, "should be a valid location")
	invalidLocation1 := customer{Latitude: 91, Longitude: 181,}
	invalidLocation2 := customer{Latitude: -91, Longitude: -181,}
	assertEqual(t, locValidate(invalidLocation1), false, "should be an invalid location")
	assertEqual(t, locValidate(invalidLocation2), false, "should be an invalid location")

	assertEqual(t, len(customers), 32, "inconsistent number from example json file")

	center := geo.NewPoint(OfficeLatitude, OfficeLongitude)
	nearCustomers := filterRange(customers, center, Distance, false)
	assertEqual(t, nearCustomers[0].ID, 12, "not correct id in the first one order")

	nearSorted := sortByID(nearCustomers)
	assertEqual(t, nearSorted[0].ID, 4, "smallest matched id should be 4")
}

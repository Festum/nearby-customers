package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bradfitz/slice"
	geo "github.com/kellydunn/golang-geo"
)

const FilePath = "./customers.json"
const OfficeLatitude = 53.339428
const OfficeLongitude = -6.257664
const Distance = 100

type customer struct {
	Latitude  float64 `json:",string,omitempty"`
	Longitude float64 `json:",string,omitempty"`
	ID        int     `json:"user_id"`
	Name      string
}

func locValidate(c customer) bool{ //Simple Validation for Latitude & Longitude
	if c.Latitude <=90.0 && c.Latitude >= -90.0 && c.Longitude <= 180.0 && c.Longitude >= -180.0 {
		return true
	}
	return false
}

func readLines(path string) ([]customer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var customers []customer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var c customer
		json.Unmarshal([]byte(scanner.Text()), &c)
		if locValidate(c){
			customers = append(customers, c)
		}
	}
	return customers, scanner.Err()
}

func filterRange(c []customer, o *geo.Point, distance float64, filter bool) []customer {
	var near []customer
	for i := 0; i < len(c); i++ {
		dist := o.GreatCircleDistance(geo.NewPoint(c[i].Latitude, c[i].Longitude))
		// log.Printf("great circle distance: %d\n", dist)
		if dist <= distance {
			if filter {
				c[i].Latitude = 0
				c[i].Longitude = 0
			}
			near = append(near, c[i])
		}
	}
	return near
}

func sortByID(c []customer) []customer {
	slice.Sort(c[:], func(i, j int) bool {
		return c[i].ID < c[j].ID
	})
	return c
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessing index")
	readme, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(readme))
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	log.Println("Invoking API: generate-plan")
	customers, err := readLines(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	near := sortByID(filterRange(customers, geo.NewPoint(OfficeLatitude, OfficeLongitude), Distance, true))
	nearJson, err := json.Marshal(near)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(nearJson)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/nearcustomers", getCustomers)
	port := ":8080"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

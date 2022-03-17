package main

import (
	"encoding/json"
	"fmt"
)

var customerList = []consumedEntry{
	{Name: "Annika", Candy: "Geisha", Eaten: 100},
	{Name: "Jonas", Candy: "Geisha", Eaten: 200},
	{Name: "Jonas", Candy: "Kexchoklad", Eaten: 100},
	{Name: "Aadya", Candy: "Nötchoklad", Eaten: 2},
	{Name: "Jonas", Candy: "Nötchoklad", Eaten: 3},
	{Name: "Jane", Candy: "Nötchoklad", Eaten: 17},
	{Name: "Annika", Candy: "Geisha", Eaten: 100},
	{Name: "Jonas", Candy: "Geisha", Eaten: 700},
	{Name: "Jane", Candy: "Nötchoklad", Eaten: 4},
	{Name: "Aadya", Candy: "Center", Eaten: 7},
	{Name: "Jonas", Candy: "Geisha", Eaten: 900},
	{Name: "Jane", Candy: "Nötchoklad", Eaten: 1},
	{Name: "Jonas", Candy: "Kexchoklad", Eaten: 12},
	{Name: "Jonas", Candy: "Plopp", Eaten: 40},
	{Name: "Jonas", Candy: "Center", Eaten: 27},
	{Name: "Aadya", Candy: "Center", Eaten: 2},
	{Name: "Annika", Candy: "Center", Eaten: 8},
}

func main() {
	customerPreferences, err := json.Marshal(setPreferredCandyForEachCustomer(mapConsumersData(customerList)))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(customerPreferences))
}

// Creates a map of consumer names and a slice of candies summed by candy type for each individual consumer
func mapConsumersData(customerList []consumedEntry) map[string][]*candy {
	uniqueConsumersData := make(map[string][]*candy)
	uniqueEvaluation := make(map[string]bool)

	for _, customer := range customerList {
		if _, ok := uniqueEvaluation[customer.Name]; !ok {
			uniqueEvaluation[customer.Name] = true

			uniqueConsumersData[customer.Name] = []*candy{{Candy: customer.Candy, Eaten: customer.Eaten}}
		} else {
			uniqueConsumersData[customer.Name] = mapCandies(uniqueConsumersData[customer.Name],
				candy{Candy: customer.Candy, Eaten: customer.Eaten})
		}
	}

	return uniqueConsumersData
}

// Map the amount of individual candies a consumer eats
func mapCandies(candies []*candy, newCandy candy) []*candy {
	var found bool

	for _, candyC := range candies {
		if candyC.Candy == newCandy.Candy {
			candyC.Eaten += newCandy.Eaten

			found = true
			break
		}
	}

	if !found {
		candies = append(candies, &newCandy)
	}

	return candies
}

func findPreferredCandy(candies []*candy) string {
	preferredCandy := candies[0]

	for _, c := range candies {
		if preferredCandy.Eaten < c.Eaten {
			preferredCandy = c
		}
	}

	return preferredCandy.Candy
}

func findTotalAmount(candies []*candy) int {
	var totalCandies int

	for _, c := range candies {
		totalCandies += c.Eaten
	}

	return totalCandies
}

// Map favourite snack adn teh total amount of sweets eaten for each
func setPreferredCandyForEachCustomer(unique map[string][]*candy) []customerPreferences {
	var customers []customerPreferences

	for customerName, candiesEaten := range unique {
		preferredCandy := findPreferredCandy(candiesEaten)
		totalCandies := findTotalAmount(candiesEaten)
		customers = append(customers, customerPreferences{Name: customerName, Candy: preferredCandy, Eaten: totalCandies})
	}

	return customers
}

type consumedEntry struct {
	Name  string `json:"name"`
	Candy string `json:"candy"`
	Eaten int    `json:"eaten"`
}

type candy struct {
	Candy string `json:"candy"`
	Eaten int    `json:"eaten"`
}

type customerPreferences struct {
	Name  string `json:"name"`
	Candy string `json:"favouriteSnack"`
	Eaten int    `json:"totalSnacks"`
}

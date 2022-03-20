package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type customerEntry struct {
	Name  string `json:"name"`
	Candy string `json:"candy"`
	Eaten int    `json:"eaten"`
}

var customerList = []customerEntry{
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
	customerPreferences, err := json.Marshal(setPreferredCandyForEachCustomer(mapCustomerData(customerList)))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(customerPreferences))
}

type candy struct {
	Candy string `json:"candy"`
	Eaten int    `json:"eaten"`
}

// Map customers to types of candies eaten
func mapCustomerData(customerList []customerEntry) map[string][]candy {
	uniqueCustomerData := make(map[string][]candy)

	for _, customer := range customerList {
		if candies, ok := uniqueCustomerData[customer.Name]; !ok {
			uniqueCustomerData[customer.Name] = []candy{{Candy: customer.Candy, Eaten: customer.Eaten}}
		} else {
			uniqueCustomerData[customer.Name] = addCandy(candies, candy{Candy: customer.Candy, Eaten: customer.Eaten})
		}
	}

	return uniqueCustomerData
}

// Return the sum of each type of candy by adding the newCandy amount
func addCandy(candies []candy, newCandy candy) []candy {
	var found bool

	for index, c := range candies {
		if c.Candy == newCandy.Candy {
			c.Eaten += newCandy.Eaten
			candies[index] = c

			found = true
			break
		}
	}

	if !found {
		candies = append(candies, newCandy)
	}

	return candies
}

// Search for the favourite candy name for a customer
func findPreferredCandy(candies []candy) string {
	preferredCandy := candies[0]

	for _, c := range candies {
		if preferredCandy.Eaten < c.Eaten {
			preferredCandy = c
		}
	}

	return preferredCandy.Candy
}

// Sum all the candies eaten by customer
func findTotalAmount(candies []candy) (totalCandies int) {
	for _, c := range candies {
		totalCandies += c.Eaten
	}

	return
}

type customerPreferences struct {
	Name  string `json:"name"`
	Candy string `json:"favouriteSnack"`
	Eaten int    `json:"totalSnacks"`
}

// Map favourite snack and the total amount of sweets eaten for each customer
func setPreferredCandyForEachCustomer(unique map[string][]candy) (customers []customerPreferences) {
	for customerName, candiesEaten := range unique {
		preferredCandy := findPreferredCandy(candiesEaten)
		totalCandies := findTotalAmount(candiesEaten)
		customers = append(customers, customerPreferences{Name: customerName, Candy: preferredCandy, Eaten: totalCandies})
	}

	sort.Sort(customerPreferencesDescending(customers))

	return
}

type customerPreferencesDescending []customerPreferences

func (c customerPreferencesDescending) Len() int           { return len(c) }
func (c customerPreferencesDescending) Less(i, j int) bool { return c[i].Eaten > c[j].Eaten }
func (c customerPreferencesDescending) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

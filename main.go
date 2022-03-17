package main

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

}

type consumedEntry struct {
	Name  string `json:"name"`
	Candy string `json:"candy"`
	Eaten int    `json:"eaten"`
}

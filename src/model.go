package main

/*
This is a simple model file defining our structs.
You have a Gopher can have many Holes:

Gopher (1) - (*) Holes

*/

type Gopher struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Born    string `json:"born"`
	Details string `json:"details"`
	Holes   []Hole `json:"holes"`
}

type Hole struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

type Repository interface {
	gopherManager(offset int, limit int) ([]Gopher, error)
}

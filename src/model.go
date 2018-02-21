package main

type Gopher struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Born    string `json:"born"`
	Details string `json:"details"`
	Holes   []Hole `json:"holes"`
}

type Gophers []Gopher

type Hole struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

type Holes []Hole

type Repository interface {
	gopherManager(offset int, limit int) (Gophers, error)
}

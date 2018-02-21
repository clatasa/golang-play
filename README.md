# golang-play

This is an example project of a very simple RESTful API in golang.

The domain is really simple:

A Gopher has many Holes. 
aka..

Gopher (1) - (* ) Holes

See model.go for Gopher and Hole attributes.

It has the following endpoints:

	/gophers GET 
	/gophers POST 

To compile and run this project try:

> go run *.go


Here's an outline of what each file does:

- main.go: Starts the server and configures logging
- router.go: configures routes and binds them to methods in handlers.go
- handlers.go: this is where requests are handled
- model.go: definition of the Structs that make up the domain
- gopherRepo.go: where the actual SQL work is done to pull the data from the DB or create it


The DB DDL is:

create table gopher (
  id BIGSERIAL,
  name TEXT,
  born TIMESTAMP,
  details JSONB,
  PRIMARY KEY (id)
);

create table hole (
  id BIGSERIAL ,
  name TEXT,
  created TIMESTAMP,
  gopher_id BIGINT REFERENCES gopher(id),
  PRIMARY KEY (id)
);

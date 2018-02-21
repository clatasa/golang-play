package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

/*
	Here we connect to the DB and do the work of either fetching gophers and their holes or creating them.

	Example query: SELECT g.id, g.name, g.born, g.details, h.id as hid, h.name as holename, created from gopher g left outer join holes h on g.id=h.gopher_id order by id offset 0 limit 50

 id |   name   |            born            |               details             | hid | holename |          created
----+----------+----------------------------+-----------------------------------+-----+----------+----------------------------
  1 | biguy    | 2018-02-11 16:53:45.356687 | {"isfast": "yes", "isbad": "yes"} |     |          |
  2 | holeking | 2018-02-11 21:26:53.32964  | {"isfast": "yes", "isbad": "yes"} |   1 | bigun    | 2018-02-19 20:40:49.091865
  2 | holeking | 2018-02-11 21:26:53.32964  | {"isfast": "yes", "isbad": "yes"} |   2 | bigun2   | 2018-02-19 20:41:01.224199
  3 | slowpoke | 2018-02-11 21:37:46.958885 | {"isfast": "no", "isbad": "yes"}  |     |          |

*/

/*
	This struct is required to deal with getting Gophers that have no Holes
*/
type SqlHole struct {
	Id      sql.NullInt64
	Name    sql.NullString
	Created sql.NullString
}

const (
	DB_USER     = "carlolatasa"
	DB_PASSWORD = ""
	DB_NAME     = "gopherholes"
	DB_LIMIT    = 50
	DB_SERVER   = "localhost"
)

var DB_CONNECT = fmt.Sprintf("postgres://%s@%s/%s?sslmode=disable", DB_USER, DB_SERVER, DB_NAME)

type dbGopherManager struct{}

func (gopherMgr *dbGopherManager) findGophers(offset int, limit int, gopherName string) []Gopher {

	//Connect and execute the query...
	db, err := sql.Open("postgres", DB_CONNECT)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queryString := fmt.Sprintf(
		"SELECT g.id, g.name, g.born, g.details, h.id as hid, h.name as holename, created from gopher g left outer join holes h on g.id=h.gopher_id ")

	if gopherName != "" {
		queryString = queryString + fmt.Sprintf("where g.name ilike '%s' ", gopherName)
	}

	queryString = queryString + fmt.Sprintf("order by id offset %d limit %d ", offset, limit)

	Info.Println(queryString)

	rows, err := db.Query(queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//Build your Structs from the result...
	var gophers []Gopher

	for rows.Next() {

		gopher := new(Gopher)
		sqlHole := new(SqlHole)

		err := rows.Scan(&gopher.Id, &gopher.Name, &gopher.Born, &gopher.Details, &sqlHole.Id, &sqlHole.Name, &sqlHole.Created)
		if err != nil {
			log.Fatal(err)
			return gophers
		}

		if sqlHole.Id.Valid {
			hole := new(Hole)
			hole.Id = int(sqlHole.Id.Int64)
			hole.Name = sqlHole.Name.String
			hole.Created = sqlHole.Created.String

			//Ugly but you have to figure out if your appending your children to a new parent or...
			if len(gophers) == 0 || gophers[len(gophers)-1].Id != gopher.Id {

				gopher.Holes = append(gopher.Holes, *hole)
			} else {
				//... the existing parent
				gophers[len(gophers)-1].Holes = append(gophers[len(gophers)-1].Holes, *hole)
			}
		}

		//only append of the row in a different parent from the previous....
		if len(gophers) == 0 || gophers[len(gophers)-1].Id != gopher.Id {
			gophers = append(gophers, *gopher)
		}

	}
	return gophers
}

func (gopherMgr *dbGopherManager) createGopher(gopher Gopher) int {

	db, err := sql.Open("postgres", DB_CONNECT)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	lastInsertId := -1
	err = db.QueryRow("insert into gopher (name,born,details) values($1, localtimestamp, $2) RETURNING id", gopher.Name, gopher.Details).Scan(&lastInsertId)

	//TODO: add support for creating holes...
	if err != nil {
		log.Fatal(err)
	}
	return lastInsertId
}

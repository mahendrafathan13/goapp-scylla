package scylla

import (
	"log"

	"github.com/gocql/gocql"
)

func SelectQuery(session *gocql.Session) {
	log.Printf("Displaying Results:\n")
	q := session.Query("SELECT * FROM ws_prediction_no_invoice")
	var firstName, lastName, address, pictureLocation string
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			log.Printf("select catalog.mutant %+v\n", err)
		}
	}()
	for it.Scan(&firstName, &lastName, &address, &pictureLocation) {
		log.Printf("\t" + firstName + " " + lastName + ", " + address + ", " + pictureLocation)
	}
}

func InsertQuery(session *gocql.Session) {
	log.Printf("Inserting Mike")
	if err := session.Query("INSERT INTO ws_prediction_no_invoice (user_id, json_prediction,) VALUES (123, 'coba')").Exec(); err != nil {
		log.Printf("select catalog.mutant %+v\n", err)
	}
}

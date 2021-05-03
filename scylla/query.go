package scylla

import (
	"log"

	"github.com/gocql/gocql"
)

func SelectQuery(session *gocql.Session) {
	log.Printf("Displaying Results:\n")
	q := session.Query("SELECT user_id, json_prediction FROM ws_prediction_no_invoice")
	var userID int64
	var jsonPrediction string
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			log.Printf("select ws_prediction_no_invoice: %+v\n", err)
		}
	}()
	for it.Scan(&userID, &jsonPrediction) {
		log.Printf("%d, %s\n", userID, jsonPrediction)
	}
}

func InsertQuery(session *gocql.Session) {
	log.Printf("Inserting New Data")
	if err := session.Query("INSERT INTO ws_prediction_no_invoice (user_id, json_prediction) VALUES (124, 'coba')").Exec(); err != nil {
		log.Printf("insert catalog.mutant: %+v\n", err)
	}
}

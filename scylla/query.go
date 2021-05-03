package scylla

import (
	"log"

	"github.com/gocql/gocql"
)

func SelectQuery(session *gocql.Session) {
	log.Printf("Displaying Results:\n")
	q := session.Query("SELECT first_name,last_name,address,picture_location FROM mutant_data")
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
	if err := session.Query("INSERT INTO mutant_data (first_name,last_name,address,picture_location) VALUES ('Mike','Tyson','1515 Main St', 'http://www.facebook.com/mtyson')").Exec(); err != nil {
		log.Printf("select catalog.mutant %+v\n", err)
	}
}

func DeleteQuery(session *gocql.Session) {
	log.Printf("Deleting Mike")
	if err := session.Query("DELETE FROM mutant_data WHERE first_name = 'Mike' and last_name = 'Tyson'").Exec(); err != nil {
		log.Printf("select catalog.mutant %+v\n", err)
	}
}

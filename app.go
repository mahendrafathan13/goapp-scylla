package main

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/tokopedia/goapp-scylla/scylla"
)

func main() {
	dsHosts := []string{"localhost:9042"}
	cluster := scylla.CreateCluster(gocql.Quorum, "tokopedia_topbot_ds", dsHosts...)
	session, err := gocql.NewSession(*cluster)
	if err != nil || session == nil {
		log.Printf("Failed to start cluster scylla db: %v\n", err)
		return
	}
	defer session.Close()

	scylla.SelectQuery(session)
	scylla.InsertQuery(session)
	scylla.SelectQuery(session)
}

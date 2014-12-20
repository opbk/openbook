package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/lib/pq/hstore"
	"log"
)

var connection *sql.DB

func InitDbConnection(driver, url string) {
	var err error
	connection, err = sql.Open(driver, url)
	if err != nil {
		log.Fatalf("Error while connection to db: %s", err)
	}
}

func Connection() *sql.DB {
	return connection
}

func HstoreToMap(hs hstore.Hstore) map[string]string {
	mp := make(map[string]string)
	for k, v := range hs.Map {
		if v.Valid {
			mp[k] = v.String
		}
	}

	return mp
}

func MapToHstore(mp map[string]string) hstore.Hstore {
	hs := hstore.Hstore{make(map[string]sql.NullString)}
	for k, v := range mp {
		hs.Map[k] = sql.NullString{v, true}
	}

	return hs
}

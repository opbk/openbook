package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"github.com/lib/pq/hstore"
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

func StringToArray(input string) []int64 {
	input = strings.Replace(input, "{", "[", -1)
	input = strings.Replace(input, "}", "]", -1)
	var result []int64
	json.Unmarshal([]byte(input), &result)
	return result
}

type RowScanner interface {
	Scan(dest ...interface{}) error
}

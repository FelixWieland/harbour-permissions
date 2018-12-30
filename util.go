package harbourpermissions

import (
	"database/sql"
	"fmt"
	"log"
)

func rowExists(query string) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Printf(query)
		log.Fatalf("error checking if row exists %v", err)
	}
	return exists
}

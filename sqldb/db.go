package sqldb

import (
	"database/sql"

	"github.com/masin556/wedding-V2-server/env"
)

var (
	sqlDb *sql.DB
)

func SetDb(db *sql.DB) {
	sqlDb = db
	if env.UseGuestbook {
		err := initializeGuestbookTable()
		if err != nil {
			panic(err)
		}
	}
	if env.UseAttendance {
		err := initializeAttendanceTable()
		if err != nil {
			panic(err)
		}
	}
}

func GetDb() *sql.DB {
	return sqlDb
}

package main

import (
	"database/sql"
	"net/http"

	"github.com/masin556/wedding-V2-server/env"
	"github.com/masin556/wedding-V2-server/httphandler"
	"github.com/masin556/wedding-V2-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()
	if env.UseGuestbook {
		mux.Handle("/api/guestbook", new(httphandler.GuestbookHandler))
	}
	if env.UseAttendance {
		mux.Handle("/api/attendance", new(httphandler.AttendanceHandler))
	}

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{env.AllowOrigin},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut},
		AllowCredentials: true,
	})

	handler := corHandler.Handler(mux)

	http.ListenAndServe(":8080", handler)
}

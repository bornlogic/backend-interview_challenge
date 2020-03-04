package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user,
		password,
		dbname,
	)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {
	a.Router.HandleFunc("/api/members", GetMembers)
	a.Router.HandleFunc("/api/members/invite", InviteMember)
	a.Router.HandleFunc("/api/members/authenticate", Authenticate)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

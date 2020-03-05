package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	services "github.com/marciusvinicius/Interview-Backend-Code-Challenge/api/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	clientOptions := options.Client().ApplyURI(
		"mongodb://localhost:27017"
	)

	client, err := mongo.Connect(
		context.TODO(),
		clientOptions
	)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {
	a.Router.HandleFunc("/api/members", services.GetMembers).Methods("GET")
	a.Router.HandleFunc("/api/members/invite", services.InviteMember).Methods("GET")
	a.Router.HandleFunc("/api/members/authenticate", services.Authenticate).Methods("POST")
	a.Router.HandleFunc("/api/whereistriangular", services.WhereIsTriangular).Methods("POST")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

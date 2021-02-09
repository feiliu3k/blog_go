package main

import (
	"database/sql"
	"net/http"

	"github.com/feiliu3k/goblog/app/http/middlewares"
	"github.com/feiliu3k/goblog/bootstrap"

	"github.com/gorilla/mux"
)

var db *sql.DB
var router *mux.Router

func main() {

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}

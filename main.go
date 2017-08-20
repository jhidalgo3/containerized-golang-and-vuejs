package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhidalgo3/containerized-golang-and-vuejs/model"
	"github.com/jhidalgo3/containerized-golang-and-vuejs/routes"
)

func getRouter() *mux.Router {
	return routes.NewRouter()
}

func main() {
	//connect := initRedis()
	//defer connect.Close()
	context := model.MongoSetup()
	defer context.Close()
	router := getRouter()
	log.Fatal(http.ListenAndServe(":3001", router))
}

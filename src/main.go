package main

import (
	"log"
	"net/http"
	"oidc_project/repository"
	"oidc_project/services"
	"oidc_project/transport"

	"github.com/gorilla/mux"

	_ "github.com/coreos/go-oidc/v3/oidc"
	_ "golang.org/x/oauth2"
)

func main() {
	repo := repository.NewMainRepository()
	serv := services.NewMainService(*repo)
	controller := transport.NewMainController(*serv)

	router := mux.NewRouter()

	router.HandleFunc("/", controller.GetData).Methods("GET")
	router.HandleFunc("/", controller.WriteData).Methods("POST")
	router.HandleFunc("/", controller.CleanData).Methods("DELETE")
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

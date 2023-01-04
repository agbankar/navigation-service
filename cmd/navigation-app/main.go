package main

import (
	v1 "github.com/agbankar/navigation-service/cmd/navigation-app/apis/v1"
	"github.com/agbankar/navigation-service/internal/model"
	"github.com/agbankar/navigation-service/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

func main() {
	visitorService := service.VisitorService{
		PageVisits: make(map[string]model.PageDetails),
		Lock:       &sync.RWMutex{},
	}
	controller := v1.NewVisitorController(visitorService)
	router := mux.NewRouter()

	router.HandleFunc("/visit", controller.Visit).Methods("POST")
	router.HandleFunc("/info", controller.GetUniqueVisits).Methods("GET")

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

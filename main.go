package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodeGetPersonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	personId, err := strconv.Atoi(vars["person_id"])
	if err != nil {
		return nil, fmt.Errorf("Invalid id: %v", err)
	}
	return getRequest{PersonId: personId}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	db, err := NewDB()
	if err != nil {
		log.Fatalf("Error in initializing database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	svc := NewService(db)

	getPersonHandler := httptransport.NewServer(
		makeGetPersonEndpoint(svc),
		decodeGetPersonRequest,
		encodeResponse,
	)

	r.Methods("GET").Path("/person/{person_id}/info").Handler(getPersonHandler)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

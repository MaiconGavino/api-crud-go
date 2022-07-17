package main

import (
	"encoding/json"
	"fmt"
	"github.com/maicongavino/api-crud-go/domain"
	"github.com/maicongavino/api-crud-go/domain/person"
	"net/http"
	"strings"
)

func main() {
	personServe, err := person.NewService("person.json")
	if err != nil {
		fmt.Printf("Error trying to create person server")
		return
	}

	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			//struct pessoa

			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("Error trying to decode body. Body should be a json. Error: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
			}
			if person.ID <= 0 {
				http.Error(w, "Error trying to create person. ID should be a positive integer", http.StatusBadRequest)
				http.Error(w, "Error trying to create person. ID should be a positive integer", http.StatusBadRequest)
			}

			//Criar pessoa
			err = personServe.Create(person)
			if err != nil {
				fmt.Printf("Error trying to create person: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		}
		if r.Method == "GET" {

			path := strings.TrimPrefix(r.URL.Path, "person")
			if path == "" {
				//Person List All
				people := personServe.List()
				err := json.NewEncoder(w).Encode(people)
				if err != nil {
					http.Error(w, "Error trying to list people", http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
			}
			//person list person id
		}
		http.Error(w, "Not Implemented", http.StatusInternalServerError)
	})
	http.ListenAndServe("8080", nil)
}

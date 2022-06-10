package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"technocar/dto"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// dummy database
var cars []dto.Car

// w as response
// r as request
func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliaction/json")

	if len(cars) < 1 {
		http.Error(w, "No car is available", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cars)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliaction/json")

	var newCar dto.Car
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		http.Error(w, "Error createCar input", http.StatusBadRequest)
		return
	}

	newCar.ID = uuid.NewString()
	newCar.CalculatePrice()

	cars = append(cars, newCar)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

func updateCarById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliaction/json")

	var updatedCar dto.Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil {
		http.Error(w, "Error createCar input", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	inputCarId := params["id"]

	for index, item := range cars {
		if item.ID == inputCarId {
			cars = append(cars[:index], cars[index+1:]...)

			updatedCar.ID = item.ID
			updatedCar.CalculatePrice()
			cars = append(cars, updatedCar)

			json.NewEncoder(w).Encode(updatedCar)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

func deleteCarById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliaction/json")

	params := mux.Vars(r)
	inputCarId := params["id"]

	for index, item := range cars {
		if item.ID == inputCarId {
			cars = append(cars[:index], cars[index+1:]...)

			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

func main() {

	// init router
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/cars", getCars).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/cars", createCar).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/cars/{id}", updateCarById).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/cars/{id}", deleteCarById).Methods(http.MethodDelete)

	fmt.Println("Server is up....")

	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8001"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}

// https://docs.google.com/document/d/1wPuJijpPJYUtZfsJzXq9kAPzalxQNMjUk5ItrBKlc-o
// https://bit.ly/TechnoScapeDW2022-Day2
// https://docs.google.com/document/d/1wbtD9Zxd8U3VUAMIKV1rU-pSK6vGN2NJuV4cVQJSYL4/

/*
Kak apakah github action yang menjalankan beberapa unit test atau integration test termasuk sandbox?
Maaf OOT kak, kak iqi kalau buat test lebih ke unit testing atau integration test ?
*/

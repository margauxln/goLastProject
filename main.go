package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/spots", getAllSpots).Methods("GET")
	router.HandleFunc("/spots/{id}", getOneSpot).Methods("GET")
	router.HandleFunc("/spot", createSpot).Methods("POST")
	router.HandleFunc("/spots/{id}",updateSpot).Methods("PATCH")
	router.HandleFunc("/spots/{id}",deleteSpot).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}

type spot struct {
	ID        string   `json:"ID"`
	Title     string   `json:"Title"`
	Address   string   `json:"Address"`
	Level     int      `json:"Level"`
	SurfBreak []string `json:"SurfBreak"`
	Photo     string   `json:"Photo"`
}

type allSpots []spot

var spots = allSpots{
	{
		ID:        "1",
		Title:     "Pipeline",
		Address:   "Pipeline, Oahu, Hawaii",
		Level:     4,
		SurfBreak: []string{"Reef Break"},
		Photo:     "https://dl.airtable.com/ZuXJZ2NnTF40kCdBfTld_thomas-ashlock-64485-unsplash.jpg",
	},
	{
		ID:        "2",
		Title:     "Skeleton Bay",
		Address:   "Skeleton Bay, Namibia",
		Level:     5,
		SurfBreak: []string{"Point Break"},
		Photo:     "https://dl.airtable.com/YzqA020RRLaTyAZAta9g_brandon-compagne-308937-unsplash.jpg",
	},
	{
		ID:        "3",
		Title:     "Superbank",
		Address:   "Superbank, Gold Coast, Australia",
		Level:     4,
		SurfBreak: []string{"Point Break"},
		Photo:     "https://dl.airtable.com/I4E4xZeQbO2g814udQDm_jeremy-bishop-80371-unsplash.jpg",
	},
}

func getAllSpots(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(spots)
}

func getOneSpot(w http.ResponseWriter, r *http.Request) {
	spotID := mux.Vars(r)["id"]

	for _, singleSpot := range spots {
		if singleSpot.ID == spotID {
			json.NewEncoder(w).Encode(singleSpot)
		}
	}
}

func createSpot(w http.ResponseWriter, r *http.Request) {
	var newSpot spot
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newSpot)
	spots = append(spots, newSpot)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newSpot)
}

func updateSpot(w http.ResponseWriter, r *http.Request) {
	spotID := mux.Vars(r)["id"]
	var updatedSpot spot

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data in order to update")
	}
	json.Unmarshal(reqBody, &updatedSpot)

	for i, singleSpot := range spots {
		if singleSpot.ID == spotID {
			singleSpot.Title = updatedSpot.Title
			singleSpot.Address = updatedSpot.Address
			singleSpot.Level = updatedSpot.Level
			singleSpot.SurfBreak = updatedSpot.SurfBreak
			singleSpot.Photo = updatedSpot.Photo
			updatedSpots := append(spots[:i], singleSpot)
			spots = append(updatedSpots, spots[i+1:]...)
			json.NewEncoder(w).Encode(singleSpot)
		}
	}
}

func deleteSpot(w http.ResponseWriter, r *http.Request) {
	spotID := mux.Vars(r)["id"]

	for i, singleSpot := range spots {
		if singleSpot.ID == spotID {
			spots = append(spots[:i], spots[i+1:]...)
			fmt.Fprintf(w, "The spot with ID %v has been deleted successfully", spotID)
		}
	}
}
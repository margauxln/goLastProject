package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

type Spot struct {
	ID        string `json:"ID"`
	Title     string `json:"Title"`
	Address   string `json:"Address"`
	Level     int    `json:"Level"`
	Photo     string `json:"Photo"`
	SurfBreak string `json:"SurfBreak"`
}

// var db *sql.DB
// var err error

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/spots", getAllSpots).Methods("GET")
	router.HandleFunc("/spots/{id}", getOneSpot).Methods("GET")
	router.HandleFunc("/spot", createSpot).Methods("POST")
	router.HandleFunc("/spots/{id}", updateSpot).Methods("PATCH")
	router.HandleFunc("/spots/{id}", deleteSpot).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}
func getAllSpots(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	ErrorCheck(err)
	defer db.Close()
	PingDB(db)

	w.Header().Set("Content-Type", "application/json")
	var allSpots []Spot
	result, err := db.Query("SELECT * from `SurfBase`.`Spot`")
	if err != nil {
		fmt.Fprintf(w, "Erreur")
		//panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var spot Spot
		err := result.Scan(&spot.ID, &spot.Title, &spot.Address, &spot.Level, &spot.Photo, &spot.SurfBreak)
		if err != nil {
			panic(err.Error())
		}
		allSpots = append(allSpots, spot)
	}
	json.NewEncoder(w).Encode(allSpots)
}

func getOneSpot(w http.ResponseWriter, r *http.Request) {
	spotID := mux.Vars(r)["id"]

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	ErrorCheck(err)

	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT * from `SurfBase`.`Spot` WHERE `id` = '" + spotID + "'")
	if err != nil {
		fmt.Fprintf(w, "Erreur")
		//panic(err.Error())
	}

	for result.Next() {
		var singleSpot Spot
		err := result.Scan(&singleSpot.ID, &singleSpot.Title, &singleSpot.Address, &singleSpot.Level, &singleSpot.Photo, &singleSpot.SurfBreak)
		if err != nil {
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(singleSpot)
	}

	defer result.Close()
	defer db.Close()
	PingDB(db)
}

func createSpot(w http.ResponseWriter, r *http.Request) {
	var newSpot Spot

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	ErrorCheck(err)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newSpot)

	insert, err := db.Query("INSERT INTO `Spot` VALUES ('" + newSpot.ID + "', '" + newSpot.Title + "', '" + newSpot.Address + "', '" + strconv.Itoa(newSpot.Level) + "', '" + newSpot.Photo + "', '" + newSpot.SurfBreak + "')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newSpot)

	defer db.Close()
	PingDB(db)
}

func updateSpot(w http.ResponseWriter, r *http.Request) {
	spotID := mux.Vars(r)["id"]
	var updatedSpot Spot

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data in order to update")
	}
	json.Unmarshal(reqBody, &updatedSpot)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	ErrorCheck(err)

	fmt.Println(updatedSpot.Photo)
	update, err := db.Query("UPDATE `Spot` SET `Title` = '" + updatedSpot.Title + "', `Address` = '" + updatedSpot.Address + "', `Level` = '" + strconv.Itoa(updatedSpot.Level) + "', `Photo` = '" + updatedSpot.Photo + "', `Surfbreak` = '" + updatedSpot.SurfBreak + "' WHERE `Spot`.`Id` = '" + spotID + "'")
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
	json.NewEncoder(w).Encode(updatedSpot)
	w.WriteHeader(http.StatusCreated)
}

func deleteSpot(w http.ResponseWriter, r *http.Request) {
	spotID := mux.Vars(r)["id"]

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	ErrorCheck(err)

	delete, err := db.Query("DELETE FROM `Spot` WHERE `Spot`.`Id` = '" + spotID + "'")
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	fmt.Fprintf(w, "The spot with ID %v has been deleted successfully", spotID)
	w.WriteHeader(http.StatusCreated)
}

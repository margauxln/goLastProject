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

//var db *sql.DB

//var db, err = sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")

//ErrorCheck(err)

// var err error

func main() {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	// ErrorCheck(err)

	// close database after all work is done
	// defer db.Close()
	// PingDB(db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/spots", getAllSpots).Methods("GET")
	//router.HandleFunc("/spots/{id}", getOneSpot).Methods("GET")
	router.HandleFunc("/spot", createSpot).Methods("POST")
	// router.HandleFunc("/spots/{id}", updateSpot).Methods("PATCH")
	// router.HandleFunc("/spots/{id}", deleteSpot).Methods("DELETE")
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

// func getOneSpot(w http.ResponseWriter, r *http.Request) {
// 	spotID := mux.Vars(r)["id"]

// 	for _, singleSpot := range spots {
// 		if singleSpot.ID == spotID {
// 			json.NewEncoder(w).Encode(singleSpot)
// 		}
// 	}
// }

func createSpot(w http.ResponseWriter, r *http.Request) {
	var newSpot Spot

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	ErrorCheck(err)
	defer db.Close()
	PingDB(db)

	insert, err := db.Query("INSERT INTO `SurfBase`.`Spot` (`id`, `Title`,`Address`,`Level`,`SurfBreak`,`Photo`) VALUES (`" + newSpot.ID + "`, `" + newSpot.Title + "`, `" + newSpot.Address + "`, `" + strconv.Itoa(newSpot.Level) + " `, `" + newSpot.SurfBreak + "`, `" + newSpot.Photo + "`)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newSpot)
	//spots = append(spots, newSpot)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newSpot)
}

// func updateSpot(w http.ResponseWriter, r *http.Request) {
// 	spotID := mux.Vars(r)["id"]
// 	var updatedSpot spot

// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data in order to update")
// 	}
// 	json.Unmarshal(reqBody, &updatedSpot)

// 	for i, singleSpot := range spots {
// 		if singleSpot.ID == spotID {
// 			singleSpot.Title = updatedSpot.Title
// 			singleSpot.Address = updatedSpot.Address
// 			singleSpot.Level = updatedSpot.Level
// 			singleSpot.SurfBreak = updatedSpot.SurfBreak
// 			singleSpot.Photo = updatedSpot.Photo
// 			updatedSpots := append(spots[:i], singleSpot)
// 			spots = append(updatedSpots, spots[i+1:]...)
// 			json.NewEncoder(w).Encode(singleSpot)
// 		}
// 	}
// }

// func deleteSpot(w http.ResponseWriter, r *http.Request) {
// 	spotID := mux.Vars(r)["id"]

// 	for i, singleSpot := range spots {
// 		if singleSpot.ID == spotID {
// 			spots = append(spots[:i], spots[i+1:]...)
// 			fmt.Fprintf(w, "The spot with ID %v has been deleted successfully", spotID)
// 		}
// 	}
// }

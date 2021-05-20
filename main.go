package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

type Spot struct {
	ID        string   `json:"ID"`
	Title     string   `json:"Title"`
	Address   string   `json:"Address"`
	Level     int      `json:"Level"`
	SurfBreak []string `json:"SurfBreak"`
	Photo     string   `json:"Photo"`
}

//var db *sql.DB

// var err error

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	ErrorCheck(err)

	// close database after all work is done
	defer db.Close()
	PingDB(db)

	// INSERT INTO DB
	// insert, err := db.Query("INSERT INTO `SurfBase`.`Spot` (`ìd`, `Title`,`Address`,`Level`,`SurfBreak`,`Photo`) VALUES ('Title test 1','Address test 1','3','Reef Break','https://upload.wikimedia.org/wikipedia/commons/thumb/5/57/Scarlett_Johansson_by_Gage_Skidmore_2_%28cropped%29_%28cropped%29.jpg/440px-Scarlett_Johansson_by_Gage_Skidmore_2_%28cropped%29_%28cropped%29.jpg')	")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/spots", getAllSpots).Methods("GET")
	// router.HandleFunc("/spots/{id}", getOneSpot).Methods("GET")
	// router.HandleFunc("/spot", createSpot).Methods("POST")
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

// type allSpots []spot

// var spots = allSpots{
// 	{
// 		ID:        "1",
// 		Title:     "Pipeline",
// 		Address:   "Pipeline, Oahu, Hawaii",
// 		Level:     4,
// 		SurfBreak: []string{"Reef Break"},
// 		Photo:     "https://dl.airtable.com/ZuXJZ2NnTF40kCdBfTld_thomas-ashlock-64485-unsplash.jpg",
// 	},
// 	{
// 		ID:        "2",
// 		Title:     "Skeleton Bay",
// 		Address:   "Skeleton Bay, Namibia",
// 		Level:     5,
// 		SurfBreak: []string{"Point Break"},
// 		Photo:     "https://dl.airtable.com/YzqA020RRLaTyAZAta9g_brandon-compagne-308937-unsplash.jpg",
// 	},
// 	{
// 		ID:        "3",
// 		Title:     "Superbank",
// 		Address:   "Superbank, Gold Coast, Australia",
// 		Level:     4,
// 		SurfBreak: []string{"Point Break"},
// 		Photo:     "https://dl.airtable.com/I4E4xZeQbO2g814udQDm_jeremy-bishop-80371-unsplash.jpg",
// 	},
// }

func getAllSpots(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/SurfBase")
	ErrorCheck(err)

	// close database after all work is done
	defer db.Close()
	PingDB(db)

	// if db == nil {
	// 	fmt.Println("db vide")
	// }
	// result, err := db.Query("SELECT * from `SurfBase`.`Spot`")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer result.Close()

	// json.NewEncoder(w).Encode(result.Next())

	// fmt.Println("Successful connection to db")
	// 	fmt.Fprintf(w, "get all spots")
	w.Header().Set("Content-Type", "application/json")
	var allSpots []Spot
	result, err := db.Query("SELECT `id`, `title` from `SurfBase`.`Spot`")
	if err != nil {
		fmt.Fprintf(w, "Erreur")
		//panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var spot Spot
		err := result.Scan(&spot.ID, &spot.Title)
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

// func createSpot(w http.ResponseWriter, r *http.Request) {
// 	var newSpot spot
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}

// 	json.Unmarshal(reqBody, &newSpot)
// 	spots = append(spots, newSpot)
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(newSpot)
// }

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

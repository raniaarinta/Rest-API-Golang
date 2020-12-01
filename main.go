package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	idUser   string `json:"iduser"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type Bike struct {
	idBike    string `json:"idbike"`
	Bikename  string `json:"bikename"`
	BikePrice int32  `json:"BikePrice"`
}
type Rent struct {
	ID         string `json:"id"`
	duration   int    `json:"duration"`
	TotalPrice int    `json:"totalPrice"`
	MyBike     *Bike  `json:"bike"`
	Myuser     *User  `json:"user"`
}

var myRent []Rent
var BikeSlice []Bike
var UserSlice []User
var bike1 Bike
var bike2 Bike
var user1 User
var user2 User

///////////////////rents
func getAllRents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myRent)

}

func getRent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range myRent {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Rent{})

}
func createRent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Rent Rent
	_ = json.NewDecoder(r.Body).Decode(&Rent)
	Rent.ID = strconv.Itoa(rand.Intn(100000000))
	myRent = append(myRent, Rent)
	json.NewEncoder(w).Encode(Rent)

}
func UpdateRent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range myRent {
		if item.ID == params["id"] {
			myRent = append(myRent[:index], myRent[index+1:]...)
			var R Rent
			_ = json.NewDecoder(r.Body).Decode(&R)
			R.ID = params["id"]
			myRent = append(myRent, R)
			json.NewEncoder(w).Encode(R)
			return
		}
	}

}
func deleteRent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range myRent {
		if item.ID == params["id"] {
			myRent = append(myRent[:index], myRent[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(myRent)

}

/////////////////////////bikes/////////////////
func getAllBikes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(BikeSlice)

}
func getBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range BikeSlice {
		if item.idBike == params["idbike"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Bike{})

}
func GetBikeName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range BikeSlice {
		if item.Bikename == params["bikename"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Bike{})
}
func createBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var B Bike
	_ = json.NewDecoder(r.Body).Decode(&B)
	B.idBike = strconv.Itoa(rand.Intn(100000000))
	BikeSlice = append(BikeSlice, B)
	json.NewEncoder(w).Encode(B)

}

func UpdateBike(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range BikeSlice {
		if item.idBike == params["idbike"] {
			BikeSlice = append(BikeSlice[:index], BikeSlice[index+1:]...)
			var B Bike
			_ = json.NewDecoder(r.Body).Decode(&B)
			B.idBike = params["idbike"]
			BikeSlice = append(BikeSlice, B)
			json.NewEncoder(w).Encode(B)
			return
		}
	}

}
func deleteBike(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range BikeSlice {
		if item.idBike == params["idbike"] {
			BikeSlice = append(BikeSlice[:index], BikeSlice[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(BikeSlice)

}

/////////////////////////user//////////////////////////
func getAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserSlice)

}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range UserSlice {
		if item.idUser == params["iduser"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})

}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var U User
	_ = json.NewDecoder(r.Body).Decode(&U)
	U.idUser = strconv.Itoa(rand.Intn(100000000))
	UserSlice = append(UserSlice, U)
	json.NewEncoder(w).Encode(U)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range UserSlice {
		if item.idUser == params["iduser"] {
			UserSlice = append(UserSlice[:index], UserSlice[index+1:]...)
			var U User
			_ = json.NewDecoder(r.Body).Decode(&U)
			U.idUser = params["iduser"]
			UserSlice = append(UserSlice, U)
			json.NewEncoder(w).Encode(U)
			return
		}
	}

}
func deleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range UserSlice {
		if item.idUser == params["iduser"] {
			UserSlice = append(UserSlice[:index], UserSlice[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(UserSlice)

}

func main() {
	//init router
	router := mux.NewRouter()
	//data
	user1 = User{idUser: "1", Username: "a", Email: "d"}
	user2 = User{idUser: "2", Username: "c", Email: "f"}
	UserSlice = append(UserSlice, user1, user2)

	bike1 = Bike{}
	bike1.idBike = "1"
	bike1.BikePrice = 300
	bike1.Bikename = "aaaaa"

	bike2 = Bike{}
	bike2.idBike = "12"
	bike2.BikePrice = 9
	bike2.Bikename = "asadsad"
	BikeSlice = append(BikeSlice, bike1, bike2)

	myRent = append(myRent, Rent{ID: "1", duration: 5, TotalPrice: 500, Myuser: &user1, MyBike: &bike1})
	myRent = append(myRent, Rent{ID: "2", duration: 5, TotalPrice: 500, Myuser: &user2, MyBike: &bike2})

	//set endpoint
	////////////////////////////user/////////////////////
	router.HandleFunc("/api/user", getAllUser).Methods("GET")
	router.HandleFunc("/api/user/{iduser}", getUser).Methods("GET")
	router.HandleFunc("/api/user", createUser).Methods("POST")
	router.HandleFunc("/api/user/{iduser}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{iduser}", deleteUser).Methods("DELETE")

	///////////////////////////bikes////////////////////////////////
	router.HandleFunc("/api/bike", getAllBikes).Methods("GET")
	router.HandleFunc("/api/bike/{idbike}", getBike).Methods("GET")
	router.HandleFunc("/api/bike/{bikename}", GetBikeName).Methods("GET")
	router.HandleFunc("/api/bike", createBike).Methods("POST")
	router.HandleFunc("/api/bike/{idbike}", UpdateBike).Methods("PUT")
	router.HandleFunc("/api/bike/{idbike}", deleteBike).Methods("DELETE")

	/////////////////////rents//////////////////////////////
	router.HandleFunc("/api/rent", getAllRents).Methods("GET")
	router.HandleFunc("/api/rent/{id}", getRent).Methods("GET")
	router.HandleFunc("/api/rent", createRent).Methods("POST")
	router.HandleFunc("/api/rent/{id}", UpdateRent).Methods("PUT")
	router.HandleFunc("/api/rent/{id}", deleteRent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}

package main

import (
	"net/http"
	"log"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
	//"database/sql"
	//"fmt"
	//"strconv"
	//m "./model"
	s "./services"
)

var CustomerServices = s.CustomerServices{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getAllAccountInfo)
	r.HandleFunc("/id={id}", searchByID)
	r.HandleFunc("/name={name}", searchByFName)
	r.HandleFunc("/lname={lname}", searchByLastname)
	r.HandleFunc("/email={email}", searchByEmail)
	r.HandleFunc("/age1={age1}&age2={age2}", searchByAge)
	r.HandleFunc("/gender={gender}", searchByGender)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getAllAccountInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	json.NewEncoder(w).Encode(CustomerServices.GetAllAccountInfo())
}

func searchByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var ID = params["id"]

	json.NewEncoder(w).Encode(CustomerServices.GetAccountInfoByID(ID))
}

func searchByFName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var name = params["name"]

	json.NewEncoder(w).Encode(CustomerServices.GetAccountInfoByFName(name))
}
func searchByLastname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var Lname = params["lname"]

	json.NewEncoder(w).Encode(CustomerServices.GetAccountInfoByLName(Lname))
}

func searchByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var Email = params["email"]

	json.NewEncoder(w).Encode(CustomerServices.GetAccountInfoByEmail(Email))
}

func searchByGender(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var Gender = params["gender"]

	json.NewEncoder(w).Encode(CustomerServices.GetAccountInfoByGender(Gender))
}

func searchByAge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var Age1 = params["age1"]
	var Age2 = params["age2"]

	json.NewEncoder(w).Encode(CustomerServices.GetAccountInfoByAge(Age1,Age2))
}
